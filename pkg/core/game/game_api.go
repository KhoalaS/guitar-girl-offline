package game

import (
	"time"

	"github.com/KhoalaS/guitar-girl-offline/pkg/core"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/main_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type GameApi struct {
	service  GameService
	timeZone string
}

func NewGameApi(service GameService) GameApi {
	return GameApi{
		service:  service,
		timeZone: "Asia/Seoul",
	}
}

func (api *GameApi) SetTimeZone(timeZone string) {
	api.timeZone = timeZone
}

func (api *GameApi) UserSave(params BaseGameRequest[user_model.UserSaveDataInfo], apiCategory string) BaseGameResponse[user_model.UserSaveRetDataInfo] {
	status, err := api.service.UserSave(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, status, err)
}

func (api *GameApi) Init(params BaseGameRequest[InitParameters], apiCategory string) BaseGameResponse[main_model.InitRetDataInfo] {
	serverList, err := api.service.Init(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, serverList, err)
}

func (api *GameApi) GetServerTime(params BaseGameRequest[GetServerTimeParams], apiCategory string) BaseGameResponse[BaseTimestamp] {
	timeStamp, err := api.service.GetServerTime(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, timeStamp, err)
}

func (api *GameApi) UserLogin(params BaseGameRequest[UserLoginParams], apiCategory string) BaseGameResponse[user_model.UserLoginRetDataInfo] {
	loginResult, serviceError := api.service.UserLogin(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, loginResult, serviceError)
}

func (api *GameApi) GetUpdateTime(params BaseGameRequest[GetUpdateTimeParams], apiCategory string) BaseGameResponse[main_model.GetUpdateTimeRetDataInfo] {
	loginResult := api.service.GetUpdateTime(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, loginResult, common_model.ErrorRetCode{})
}

func (api *GameApi) DefaultSettingList(params BaseGameRequest[DefaultSettingListParams], apiCategory string) BaseGameResponse[DefaultSettingList] {
	settingList := api.service.DefaultSettingList(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, settingList, common_model.ErrorRetCode{})
}

func (api *GameApi) GetGameDataList(params BaseGameRequest[GetGameDataListParams], apiCategory string) BaseGameResponse[map[string]any] {
	settingList, err := api.service.GetGameDataList(params.Data)

	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, settingList, err)
}

func NewBaseGameResponse[T any](functionName string, apiCategory string, timeZone string, data T, serviceError common_model.ErrorRetCode) BaseGameResponse[T] {
	now := time.Now()

	baseResponse := BaseGameResponse[T]{
		Timestamp:    getBaseTimeStamp(now, timeZone),
		Category:     apiCategory,
		FunctionName: functionName,
		EmptyValue:   struct{}{},
	}

	baseResponse.Error = serviceError
	if serviceError.Code == 0 {
		baseResponse.Data = data
	}

	return baseResponse
}

func SecondsToServerISO(t time.Time, timeZone string) string {
	loc, _ := time.LoadLocation(timeZone)
	return t.In(loc).Format("20060102150405")
}

func getBaseTimeStamp(t time.Time, timeZone string) BaseTimestamp {
	return BaseTimestamp{
		UnixSeconds:       t.Unix(),
		ServerTimeIsoDate: int64(core.MustAtoi(SecondsToServerISO(t, timeZone))),
	}
}

type BaseTimestamp struct {
	UnixSeconds       int64 `thrift:",1"`
	ServerTimeIsoDate int64 `thrift:",2"`
}

type EnvironmentData struct {
	Version     int32  `thrift:",1"`
	Environment string `thrift:",2"`
	UnknownFlag int16  `thrift:",3"`
}

type BaseGameRequest[T any] struct {
	FunctionName string          `thrift:",1"`
	Data         T               `thrift:",2"`
	Environment  EnvironmentData `thrift:",3"`
}

type BaseGameResponse[T any] struct {
	Error        common_model.ErrorRetCode `thrift:",1"`
	Timestamp    BaseTimestamp             `thrift:",2"`
	Category     string                    `thrift:",3"`
	FunctionName string                    `thrift:",4"`
	Data         T                         `thrift:",5,omitempty"`
	EmptyValue   Empty                     `thrift:",6"`
}

type Empty struct{}
