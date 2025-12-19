package game

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/KhoalaS/guitar-girl-offline/pkg/rpc"
	"github.com/rs/zerolog/log"
	"github.com/thrift-iterator/go/general"
	"github.com/thrift-iterator/go/protocol"
)

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

func (gameRpc *GameRpc) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/main/", http.StripPrefix("/main", gameRpc.mainMux))
	mux.Handle("/post/", http.StripPrefix("/post", gameRpc.postMux))
	mux.Handle("/user/", http.StripPrefix("/user", gameRpc.userMux))
	mux.Handle("/store/", http.StripPrefix("/store", gameRpc.storeMux))
}

func (gameRpc *GameRpc) registerMainMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/Request/en", gameRpc.mainRequest)

	//TODO other endpoints

	gameRpc.mainMux = mux
}

func (gameRpc *GameRpc) registerPostMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getPostTime/en", gameRpc.getPostTime)

	//TODO other endpoints

	gameRpc.postMux = mux
}

func (gameRpc *GameRpc) registerUserMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/userJoin/en", gameRpc.userJoin)

	//TODO other endpoints

	gameRpc.userMux = mux
}

func (gameRpc *GameRpc) registerStoreMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getVarietyStore/en", gameRpc.getVarietyStore)

	//TODO other endpoints

	gameRpc.storeMux = mux
}

func (gameRpc *GameRpc) mainRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 1).Send()
		return
	}

	requestDto := FormDataToRpcRequestDto(r.Form)
	log.Debug().Str("data", requestDto.TapsonicData).Send()
	generalStruct, err := rpc.ThriftDataToStruct(requestDto.TapsonicData)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 2).Send()
		return
	}
	baseRequest := ThriftStructToBaseGameRequest(generalStruct, InitParametersMapperFunc)

	log.Info().Any("baseRequest", baseRequest).Send()

	if baseRequest.FunctionName == "init" {
		response := gameRpc.api.Init(baseRequest, "main")
		responseData, _ := json.Marshal(response)
		w.Write(responseData)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (gameRpc *GameRpc) mainInit(w http.ResponseWriter, r *http.Request) {

}

func (gameRpc *GameRpc) getPostTime(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) userJoin(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) getVarietyStore(w http.ResponseWriter, r *http.Request) {
}

func ThriftStructToBaseGameRequest[T any](thriftStruct general.Struct, dataMapperFunc func(general.Struct) T) BaseGameRequest[T] {

	return BaseGameRequest[T]{
		FunctionName: thriftStruct.Get(protocol.FieldId(1)).(string),
		Data:         dataMapperFunc(thriftStruct.Get(protocol.FieldId(2)).(general.Struct)),
		Environment:  EnvironmentMapperFunc(thriftStruct.Get(protocol.FieldId(3)).(general.Struct)),
	}
}

func FormDataToRpcRequestDto(formData url.Values) RpcRequestDto {
	log.Debug().Any("formData", formData).Send()
	return RpcRequestDto{
		Call:         formData.Get("call"),
		TapsonicData: formData.Get("tapsonic_data"),
		AccessToken:  formData.Get("access_token"),
		CacheControl: formData.Get("cache_control"),
		CurrentTime:  formData.Get("current_time"),
	}
}

type RpcRequestDto struct {
	Call         string `json:"call"`
	TapsonicData string `json:"tapsonic_data"`
	AccessToken  string `json:"access_token"`
	CacheControl string `json:"cache_control"`
	CurrentTime  string `json:"current_time"`
}

func InternalErrorHandler(w http.ResponseWriter, err error) {
	log.Error().Err(err).Msg("[Game Server]: Internal server error")
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
