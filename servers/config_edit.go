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
	if err == nil {
		w.WriteHeader(200)
		return
	}

	switch err.(type) {
	case configuration.ParsingError:
		w.WriteHeader(404)
	case configuration.EditingError:
		w.WriteHeader(500)
	default:
		w.WriteHeader(500)
	}
	w.Write([]byte(err.Error()))

	return
}
