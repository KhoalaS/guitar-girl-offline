package game

import "net/http"

type GameRpc struct {
	api      GameApi
	mainMux  *http.ServeMux
	postMux  *http.ServeMux
	userMux  *http.ServeMux
	storeMux *http.ServeMux
}

func NewGameRpc(api GameApi) *GameRpc {
	gameRpc := &GameRpc{
		api: api,
	}

	gameRpc.registerMainMux()
	gameRpc.registerPostMux()
	gameRpc.registerUserMux()
	gameRpc.registerStoreMux()

	return gameRpc
}

func (rpc *GameRpc) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/main/", http.StripPrefix("/main/", rpc.mainMux))
	mux.Handle("/post/", http.StripPrefix("/post/", rpc.postMux))
	mux.Handle("/user/", http.StripPrefix("/user/", rpc.userMux))
	mux.Handle("/store/", http.StripPrefix("/store/", rpc.storeMux))
}

func (rpc *GameRpc) registerMainMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/Request/en", rpc.mainRequest)

	//TODO other endpoints

	rpc.mainMux = mux
}

func (rpc *GameRpc) registerPostMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getPostTime/en", rpc.getPostTime)

	//TODO other endpoints

	rpc.postMux = mux
}

func (rpc *GameRpc) registerUserMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/userJoin/en", rpc.userJoin)

	//TODO other endpoints

	rpc.userMux = mux
}

func (rpc *GameRpc) registerStoreMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getVarietyStore/en", rpc.getVarietyStore)

	//TODO other endpoints

	rpc.storeMux = mux
}

func (rpc *GameRpc) mainRequest(w http.ResponseWriter, r *http.Request) {
}
func (rpc *GameRpc) getPostTime(w http.ResponseWriter, r *http.Request) {
}
func (rpc *GameRpc) userJoin(w http.ResponseWriter, r *http.Request) {
}
func (rpc *GameRpc) getVarietyStore(w http.ResponseWriter, r *http.Request) {
}
