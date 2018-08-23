package servers

import (
	"net/http"
	"github.com/alohen/dynamic_configuration_manager/configuration/page_builder"
	"github.com/alohen/dynamic_configuration_manager/configuration"
)

const (
	ReadConfigPrefix   = "/read/"
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
		http.Error(w,err.Error(),server.getStatusCodeByError(err))
	}

	w.WriteHeader(200)
	w.Write(page)
}

func (server *ConfigRetrieveServer) getStatusCodeByError(err error) int {
	switch err.(type) {
	case configuration.ParsingError:
		return http.StatusInternalServerError
	case configuration.PageBuildingError:
		return http.StatusNotFound
	default:
		return http.StatusNotFound
	}
}
