package game

import (
	"encoding/json"
	"net/http"
	"net/url"

	embeds "github.com/KhoalaS/guitar-girl-offline"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/rpc"
	"github.com/KhoalaS/thrifter"
	"github.com/KhoalaS/thrifter/general"
	"github.com/KhoalaS/thrifter/protocol"
	"github.com/rs/zerolog/log"
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

	mux.HandleFunc("/Request/en/", gameRpc.mainRequest)
	mux.HandleFunc("/getUpdateTime/en/", gameRpc.getUpdateTime)
	mux.HandleFunc("/defaultSettingList/en/", gameRpc.defaultSettingList)
	mux.HandleFunc("/getGameDataList/en/", gameRpc.getGameDataList)
	mux.HandleFunc("/getServerTime/en/", gameRpc.getServerTime)
	mux.HandleFunc("/getEventRewardList/en/", gameRpc.getEventRewardList)

	//TODO other endpoints

	gameRpc.mainMux = mux
}

func (gameRpc *GameRpc) registerPostMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getPostTime/en/", gameRpc.getPostTime)

	//TODO other endpoints

	gameRpc.postMux = mux
}

func (gameRpc *GameRpc) registerUserMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/userJoin/en/", gameRpc.userJoin)
	mux.HandleFunc("/userLogin/en/", gameRpc.userLogin)
	mux.HandleFunc("/userSave/en/", gameRpc.userLogin)

	//TODO other endpoints

	gameRpc.userMux = mux
}

func (gameRpc *GameRpc) registerStoreMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getVarietyStore/en/", gameRpc.getVarietyStore)

	//TODO other endpoints

	gameRpc.storeMux = mux
}

func (gameRpc *GameRpc) mainRequest(w http.ResponseWriter, r *http.Request) {
	generalStruct, err := getGeneralStructFromRequest(r)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Err(err).Send()
		return
	}

	var responseData []byte

	functionName := GetFunctionNameFromThrift(generalStruct)
	log.Debug().Str("name", functionName).Send()

	switch functionName {
	case "init":
		baseRequest := ThriftStructToBaseGameRequest(generalStruct, 3, InitParametersMapperFunc)
		response := gameRpc.api.Init(baseRequest, "main")
		responseData, err = thrifter.Marshal(response)
		if err != nil {
			InternalErrorHandler(w, err)
			log.Debug().Int("code", 3).Err(err).Send()
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Form.Get("debug") == "true" {
		var responseStructData general.Struct
		thrifter.Unmarshal(responseData, &responseStructData)
		log.Debug().Any("debugData", responseStructData).Send()

		responseData, _ = json.Marshal(responseStructData)
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}

func (gameRpc *GameRpc) mainInit(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) getPostTime(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) userJoin(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) userLogin(w http.ResponseWriter, r *http.Request) {
	generalStruct, err := getGeneralStructFromRequest(r)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Err(err).Send()
		return
	}

	baseRequest := ThriftStructToBaseGameRequest(generalStruct, 3, UserLoginParamsMapperFunc)
	response := gameRpc.api.UserLogin(baseRequest, "user")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}
func (gameRpc *GameRpc) getVarietyStore(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) getUpdateTime(w http.ResponseWriter, r *http.Request) {
	generalStruct, err := getGeneralStructFromRequest(r)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Err(err).Send()
		return
	}

	baseRequest := ThriftStructToBaseGameRequest(generalStruct, 3, GetUpdateTimeMapperFunc)
	response := gameRpc.api.GetUpdateTime(baseRequest, "main")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}

func (gameRpc *GameRpc) defaultSettingList(w http.ResponseWriter, r *http.Request) {
	generalStruct, err := getGeneralStructFromRequest(r)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Err(err).Send()
		return
	}

	baseRequest := ThriftStructToBaseGameRequest(generalStruct, 3, DefaultSettingListMapperFunc)
	response := gameRpc.api.DefaultSettingList(baseRequest, "main")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}

func (gameRpc *GameRpc) getGameDataList(w http.ResponseWriter, r *http.Request) {
	w.Write(embeds.RawGameData)
}

func (gameRpc GameRpc) getServerTime(w http.ResponseWriter, r *http.Request) {
	generalStruct, err := getGeneralStructFromRequest(r)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Err(err).Send()
		return
	}

	baseRequest := ThriftStructToBaseGameRequest(generalStruct, 4, GetServerTimeParamsMapperFunc)
	baseRequest.Data.TimeZone = gameRpc.api.timeZone
	response := gameRpc.api.GetServerTime(baseRequest, "main")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}

func (gameRpc GameRpc) getEventRewardList(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Write([]byte("QlpoNDFBWSZTWQiziR0ACIp//f//8hts5H///wX6I6///aHuGAUP/CAdxAB/jF4gAFAFHegXWvdnRBQACDJgBMAATAAAAAJgmACYAAJiYAACYmICIJ6RNPVNMgyAMRoAADQBiBoyAYgAYgyGIaaAyaaMcJgAAAAACYRgAAAAmJhNBk0ZGAAAAAE1UQknkQ0obJpN6oMgD1AYmjQA0NNAyBiANNHqeSAe1QaAaZ6oIkqKZqeSNGmgAAaNAAANAAA0ANNAA0AAAAA1SzpunFNOlOCVx1aJclXFvi36mMRkWBURIABYgpcS2C20CwttCWYlyKm7Rbt1W2uCrVa1mLK5OGRhpTNusWzPHGylxG1rupM26aHEcU1rXXNxWxxJbo4YNDdijHDaM2aWTFZt5FnEVYWINfQylpiqpwwk3YQ0sSrHTVMpcMFVyrCla4nDLCptm21GspiFuxjAtlTTM2WICkS9pAbqwlIDAIgOjwcDBxwBaQCFACAAIIBAC4VuVbo12InEVAOc4nE6hhETsBMIFUMarsG049AZ+IqhSAzFgALfxBYsCuWmFRnEYjHTpiJZwrBi7MQixBCoAsgCyCjIgtX1BWNq64+m1K1MBVwZrZzbyRtsn4FNHa8xeL7d77/P5h5YtSrIk/cut+TZ37PrrVCdGfsRydE1eqhYr54ddcVl1MpCzAFnZQSlBdIEGQxTUJBKHYEMgkJCDIpauUtywZAkuhuYfJ/35t87mY+43kyAUsKKIQqJfmIiHaBvQv4IoZWI5S7GVUDTCaYnffrjvztmY+cbD8BiU3nZOWMMXhrjvKLirAXTWrI4GoD9U1ZS9eeYT+iwWT31VjDvoWc/RV4vdgvcnexvD4rIsGId9tZwjRQUUyiNEooqqqqhYBDSXrSCQkjiWaAzWvN4PNselYotVMHzMB/2O9R/Wjz48KPUjaMGI2jdK7cc8e6Or+WOwq5pLFTMdirMcSWpVg1PSjZW9bo7kdc8g2kt5iOCWKxTC/djtVvYXPGL2lGO2dw8dPl0d028SPC8SNPY0cO7GSRCwujgLDIgXRISMbUUCyvS9q3YOFHSVwUvAjfHuKMGhV5UHxsx044Kp1YyVcp16sJ49GEVrHvP5HdjB82Pw0eYfxOWOdHOd6j0fqHkJdVSwVdM6l7XrRn7lGObnBwjfR5tHPut/fUZ789vBngvNzoOONiWngVrRvjS72NjcE8NSzRuo7R1KM9ejmjG2R1Dsx58ffRtUnNB2lOsf8kv1R/RL10uiPUks1HRJaRlL/VHSzR60f7Mx+2jjkupxxrJfQktGY3R5MaxryUYto9WPR1o8qP/RmS2kuiS9WNqOrR3I6zwzHrZ7EctHyUvJOiPLo6IxHJJZj8hxSWsfZPXjfHlxrR5lHJGsfANSrSOiOt/8XckU4UJAIs4kdA="))
}

func (gameRpc *GameRpc) userSave(w http.ResponseWriter, r *http.Request) {
	// TODO actual saving
	response := gameRpc.api.UserSave(BaseGameRequest[user_model.UserSaveDataInfo]{}, "main")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}

func GetFunctionNameFromThrift(thriftStruct general.Struct) string {
	return thriftStruct.Get(protocol.FieldId(1)).(string)
}

func getGeneralStructFromRequest(r *http.Request) (general.Struct, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	requestDto := FormDataToRpcRequestDto(r.Form)
	generalStruct, err := rpc.ThriftDataToStruct(requestDto.TapsonicData)
	if err != nil {
		return nil, err
	}

	return generalStruct, nil
}

func ThriftStructToBaseGameRequest[T any](thriftStruct general.Struct, environmentThriftIndex int16, dataMapperFunc func(general.Struct) T) BaseGameRequest[T] {

	environmentRawStruct := thriftStruct.Get(protocol.FieldId(environmentThriftIndex))
	environment := EnvironmentMapperFunc(environmentRawStruct.(general.Struct))

	return BaseGameRequest[T]{
		FunctionName: thriftStruct.Get(protocol.FieldId(1)).(string),
		Data:         dataMapperFunc(thriftStruct.Get(protocol.FieldId(2)).(general.Struct)),
		Environment:  environment,
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
		Debug:        formData.Get("debug") == "true",
	}
}

type RpcRequestDto struct {
	Call         string `json:"call"`
	TapsonicData string `json:"tapsonic_data"`
	AccessToken  string `json:"access_token"`
	CacheControl string `json:"cache_control"`
	CurrentTime  string `json:"current_time"`
	Debug        bool   `json:"debug"`
}

func InternalErrorHandler(w http.ResponseWriter, err error) {
	log.Error().Err(err).Msg("[Game Server]: Internal server error")
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
