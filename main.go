package main

import (
	"log"
	"net/http"
	"path/filepath"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/alohen/dynamic_configuration_manager/configuration"
	"github.com/alohen/dynamic_configuration_manager/configuration/editor"
	"github.com/alohen/dynamic_configuration_manager/configuration/page_builder"
	"github.com/alohen/dynamic_configuration_manager/servers"
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

	configLoader := configuration.NewConfigLoader(absConfigDirectory)
	configEditor := editor.NewConfigEditor(configLoader)
	editingPageBuilder := page_builder.NewEditingPageBuilder(configLoader)

	retrieveServer := servers.NewConfigReadingServer(editingPageBuilder)
	editServer := servers.NewConfigEditingServer(configEditor)
	resourceServer := http.FileServer(http.Dir("assets"))

	http.Handle(resourcesDir, http.StripPrefix(resourcesDir, resourceServer))
	http.Handle(servers.ReadConfigPrefix, http.StripPrefix(servers.ReadConfigPrefix, retrieveServer))
	http.Handle(servers.EditingUrlPrefix, http.StripPrefix(servers.EditingUrlPrefix, editServer))

	log.Printf("Trying to run server from %s on %s\n", absConfigDirectory, *serverHostAndPort)
	err = http.ListenAndServe(*serverHostAndPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
