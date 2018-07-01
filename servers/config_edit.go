// @CR: file ought to be named config_editor

package servers

import (
	"io/ioutil"
	"net/http"
	"github.com/alohen/dynamic_configuration_manager/configuration/editor"
	"github.com/alohen/dynamic_configuration_manager/configuration"
)

const (
	EditingUrlPrefix = "/edit/"
)

type ConfigEditingServer struct {
	configEditor *editor.ConfigEditor
}

func NewConfigEditingServer(configEditor *editor.ConfigEditor) http.Handler {
	return &ConfigEditingServer{
		configEditor: configEditor,
	}
}

func (server *ConfigEditingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	err = server.configEditor.EditConfiguration(filePath, body)
	if err != nil {
		http.Error(w,err.Error(),server.getStatusCodeByError(err))
		return
	}

	w.WriteHeader(200)
}

func (server *ConfigEditingServer) getStatusCodeByError(err error) int {
	switch err.(type) {
	case configuration.ParsingError:
		return http.StatusNotFound
	case configuration.EditingError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
