package servers

import (
	"net/http"
	"github.com/alohen/dynamic_configuration_manager/configuration/page_builder"
)

const (
	ReadConfigPrefix   = "/read/"
	MissingConfigError = "No such example_config"
	PageBuildingError  = "Couldn't build page"
)

type ConfigRetrieveServer struct {
	editingPageBuilder *page_builder.PageBuilder
}

func NewConfigReadingServer(editingPageBuilder *page_builder.PageBuilder) http.Handler {
	return &ConfigRetrieveServer{
		editingPageBuilder: editingPageBuilder,
	}
}
func (server *ConfigRetrieveServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	page, err := server.editingPageBuilder.BuildEditingPage(filePath)
	if err != nil {

	}

	w.WriteHeader(200)
	w.Write(page)
}

