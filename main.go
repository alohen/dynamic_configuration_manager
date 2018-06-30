package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alohen/dynamic_configuration_manager/configuration"
	"github.com/alohen/dynamic_configuration_manager/servers"
	"github.com/alohen/dynamic_configuration_manager/configuration/editor"
)

const(
	resourcesDir = "/assets/"
	ConfigPath = "example_config\\configuration"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	configLoader := configuration.NewConfigLoader(cwd,ConfigPath)
	configEditor := editor.NewConfigEditor(configLoader)
	retrieveServer := servers.NewConfigReadingServer(configLoader)
	editServer := servers.NewConfigEditingServer(configEditor)
	resourceServer := http.FileServer(http.Dir("assets"))

	http.Handle(resourcesDir, http.StripPrefix(resourcesDir, resourceServer))
	http.Handle(servers.ReadConfigPrefix, http.StripPrefix(servers.ReadConfigPrefix, retrieveServer))
	http.Handle(servers.EditingUrlPrefix, http.StripPrefix(servers.EditingUrlPrefix, editServer))

	hostAndPort := "localhost:8080"
	log.Printf("Trying to run server from %s on %s\n", cwd, hostAndPort)
	err = http.ListenAndServe(hostAndPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
