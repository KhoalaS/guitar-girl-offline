package auth

import (
	"encoding/json"
	"io"
	"io/fs"
	"net/http"

	embeds "github.com/KhoalaS/guitar-girl-offline"
	"github.com/rs/zerolog/log"
)

type AuthApi struct {
	authService AuthService
}

func NewAuthApi(service AuthService) *AuthApi {
	return &AuthApi{
		authService: service,
	}
}

func (auth *AuthApi) RegisterStaticRoutes(mux *http.ServeMux) {
	staticSubDir, _ := fs.Sub(embeds.StaticFiles, "static")
	eulaFile, _ := staticSubDir.Open("eula.html")
	eulaFileBytes, _ := io.ReadAll(eulaFile)

	mux.Handle(
		"/images/",
		http.FileServerFS(staticSubDir),
	)
	mux.Handle(
		"/css/",
		http.FileServerFS(staticSubDir),
	)
	mux.Handle(
		"/js/",
		http.FileServerFS(staticSubDir),
	)
	mux.HandleFunc("/eula.html", func(w http.ResponseWriter, r *http.Request) {
		w.Write(eulaFileBytes)
	})
}

func (auth *AuthApi) RegisterAccountsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/accounts/v3/global/login", auth.loginHandler)
}

func (auth *AuthApi) loginHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		InternalErrorHandler(w, err)
		return
	}

	requestDto := LoginRequestDtoFromFormdata(r.Form)

	response := auth.authService.Login(requestDto)
	responseData, err := json.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		return
	}

	w.Write(responseData)
}

func InternalErrorHandler(w http.ResponseWriter, err error) {
	log.Error().Err(err).Msg("Internal server error")
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
