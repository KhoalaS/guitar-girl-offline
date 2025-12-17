package core

import (
	"io"
	"io/fs"
	"net/http"

	embeds "github.com/KhoalaS/guitar-girl-offline"
)

type AuthServer struct {
	api AuthApi
	mux *http.ServeMux
}

func NewServer(api AuthApi) *AuthServer {
	return &AuthServer{
		api: api,
		mux: http.NewServeMux(),
	}
}

func (server *AuthServer) Run() {
	staticSubDir, _ := fs.Sub(embeds.StaticFiles, "static")
	eulaFile, _ := staticSubDir.Open("eula.html")
	eulaFileBytes, _ := io.ReadAll(eulaFile)
	
	server.mux.Handle(
		"/images/",
		http.FileServerFS(staticSubDir),
	)
	server.mux.Handle(
		"/css/",
		http.FileServerFS(staticSubDir),
	)
	server.mux.Handle(
		"/js/",
		http.FileServerFS(staticSubDir),
	)
	server.mux.HandleFunc("/eula.html", func(w http.ResponseWriter, r *http.Request) {
		w.Write(eulaFileBytes)
	})
	http.ListenAndServe(":9999", server.mux)
}
