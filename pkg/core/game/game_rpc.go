package game

import (
	"net/http"
	"net/url"

	embeds "github.com/KhoalaS/guitar-girl-offline"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/main_model"
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
	mux.HandleFunc("/userSave/en/", gameRpc.userSave)

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
	requestDto, _ := getRequestDto(r)

	var responseData []byte

	baseRequest, _ := rpc.ThriftDataToAny[main_model.Request](requestDto.TapsonicData)
	response := gameRpc.api.Init(baseRequest, "main")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}

func (gameRpc *GameRpc) getPostTime(w http.ResponseWriter, r *http.Request) {
}
func (gameRpc *GameRpc) userJoin(w http.ResponseWriter, r *http.Request) {
	requestDto, _ := getRequestDto(r)

	var responseData []byte

	baseRequest, _ := rpc.ThriftDataToAny[user_model.UserJoin](requestDto.TapsonicData)

	response := gameRpc.api.UserJoin(baseRequest, "user")
	responseData, err := thrifter.Marshal(response)
	if err != nil {
		InternalErrorHandler(w, err)
		log.Debug().Int("code", 3).Err(err).Send()
		return
	}

	thriftBytes, _ := rpc.ThriftBytesToBz2B64(responseData)
	w.Write([]byte(thriftBytes))
}
func (gameRpc *GameRpc) userLogin(w http.ResponseWriter, r *http.Request) {
	requestDto, _ := getRequestDto(r)

	baseRequest, _ := rpc.ThriftDataToAny[user_model.UserLogin](requestDto.TapsonicData)
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
	// TODO
	w.Write([]byte("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"))
}
func (gameRpc *GameRpc) getUpdateTime(w http.ResponseWriter, r *http.Request) {
	requestDto, _ := getRequestDto(r)

	baseRequest, _ := rpc.ThriftDataToAny[main_model.GetUpdateTime](requestDto.TapsonicData)
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
	requestDto, _ := getRequestDto(r)

	baseRequest, _ := rpc.ThriftDataToAny[main_model.DefaultSettingList](requestDto.TapsonicData)
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
	requestDto, _ := getRequestDto(r)

	baseRequest, _ := rpc.ThriftDataToAny[main_model.GetServerTime](requestDto.TapsonicData)
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

func getRequestDto(r *http.Request) (RpcRequestDto, error) {
	err := r.ParseForm()
	if err != nil {
		return RpcRequestDto{}, err
	}

	requestDto := FormDataToRpcRequestDto(r.Form)
	return requestDto, nil
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
