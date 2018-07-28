package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/alohen/dynamic_configuration_manager/config_handeling"
	"github.com/alohen/dynamic_configuration_manager/servers"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	resourcesDir = "/assets/"
)

var (
	configDirectory   = kingpin.Arg("confdir", "Directory of configuration.").Required().String()
	serverHostAndPort = kingpin.Flag("server", "Server host+port to listen on.").Short('s').Default("localhost:8080").String()
)

func main() {
	kingpin.Parse()

	absConfigDirectory, err := filepath.Abs(*configDirectory)
	if err != nil {
		log.Panic(err)
	}

	configLoader := config_handeling.ConfigLoader{ConfigDirectory: absConfigDirectory}
	retrieveServer := servers.ConfigRetrieveServer{ConfigLoader: &configLoader}
	editServer := servers.ConfigEditingServer{ConfigLoader: &configLoader}
	resourceServer := http.FileServer(http.Dir("assets"))

	http.Handle(resourcesDir, http.StripPrefix(resourcesDir, resourceServer))
	http.HandleFunc(servers.ReadConfigPrefix, retrieveServer.ServeHTTP)
	http.HandleFunc(servers.EditingUrlPrefix, editServer.ServeHTTP)

	log.Printf("Trying to run server from %s on %s\n", absConfigDirectory, *serverHostAndPort)
	err = http.ListenAndServe(*serverHostAndPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
