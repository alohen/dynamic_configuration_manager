package main

import (
	"net/http"
	"dynamic_config_editor/config_handeling"
	"dynamic_config_editor/structs"
	"dynamic_config_editor/servers"
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


