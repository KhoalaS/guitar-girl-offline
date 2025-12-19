package debug

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/KhoalaS/guitar-girl-offline/pkg/rpc"
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

		dataBytes, _ := json.Marshal(data)
		w.Write(dataBytes)
	})

	return &DebugServer{
		mux: mux,
	}
}

func (server *DebugServer) Run() {
	http.ListenAndServe(":9999", server.mux)
}
