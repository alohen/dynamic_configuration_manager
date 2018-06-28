package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alohen/dynamic_configuration_manager/config_handeling"
	"github.com/alohen/dynamic_configuration_manager/servers"
)

const(
	resourcesDir = "/assets/"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	configLoader := config_handeling.ConfigLoader{WorkingDirectory: cwd}
	retrieveServer := servers.ConfigRetrieveServer{ConfigLoader: &configLoader}
	editServer := servers.ConfigEditingServer{ConfigLoader: &configLoader}
	resourceServer := http.FileServer(http.Dir("assets"))

	http.Handle(resourcesDir, http.StripPrefix(resourcesDir, resourceServer))
	http.HandleFunc(servers.ReadConfigPrefix, retrieveServer.ServeHTTP)
	http.HandleFunc(servers.EditingUrlPrefix, editServer.ServeHTTP)

	hostAndPort := "localhost:8080"
	log.Printf("Trying to run server from %s on %s\n", cwd, hostAndPort)
	err = http.ListenAndServe(hostAndPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
