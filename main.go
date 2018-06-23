package main

import (
	"net/http"
	"github.com/alohen/dynamic_configuration_manager/config_handeling"
	"github.com/alohen/dynamic_configuration_manager/structs"
	"github.com/alohen/dynamic_configuration_manager/servers"
	"log"
)

func main() {
	configLoader := config_handeling.ConfigLoader{structs.WorkingDirectory,}
	retrieveServer := servers.ConfigRetrieveServer{&configLoader}
	editServer := servers.ConfigEditingServer{&configLoader}

	http.HandleFunc(servers.ReadConfigPrefix, retrieveServer.ServeHTTP)
	http.HandleFunc(servers.EditingUrlPrefix, editServer.ServeHTTP)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}


