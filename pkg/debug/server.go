package debug

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KhoalaS/guitar-girl-offline/pkg/rpc"
	"github.com/thrift-iterator/go/general"
)

type DebugServer struct {
	mux *http.ServeMux
}

func NewDebugServer() *DebugServer {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /translate", func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()

		data, err := rpc.ThriftDataToStruct(string(bodyBytes))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataBytes, err := json.Marshal(data)

		if err != nil {
			
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(dataBytes)
	})

	return &DebugServer{
		mux: mux,
	}
}

func (server *DebugServer) Run() {
	http.ListenAndServe(":9999", server.mux)
}

type JSONMapWrapper struct {
	Map general.Map
}

func (w JSONMapWrapper) MarshalJSON() ([]byte, error) {
	tmp := make(map[string]interface{})

	for k, v := range w.Map {
		key, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("json: key must be string, got %T", k)
		}
		tmp[key] = v
	}

	return json.Marshal(tmp)
}
