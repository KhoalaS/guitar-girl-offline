package game

import (
	"time"

	"github.com/KhoalaS/guitar-girl-offline/pkg/core"
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

func (api *GameApi) Init(params BaseGameRequest[InitParameters], apiCategory string) BaseGameResponse[InitServerUrls] {
	serverList, err := api.service.Init(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, serverList, err)
}

func (api *GameApi) GetServerTime(params BaseGameRequest[GetServerTimeParams], apiCategory string) BaseGameResponse[BaseTimestamp] {
	timeStamp, err := api.service.GetServerTime(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, timeStamp, err)
}

func (api *GameApi) UserLogin(params BaseGameRequest[UserLoginParams], apiCategory string) BaseGameResponse[*UserLoginResult] {
	loginResult, serviceError := api.service.UserLogin(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, loginResult, serviceError)
}

func NewBaseGameResponse[T any](functionName string, apiCategory string, timeZone string, data T, serviceError *ServiceError) BaseGameResponse[T] {
	now := time.Now()

	baseResponse := BaseGameResponse[T]{
		Timestamp:    getBaseTimeStamp(now, timeZone),
		Category:     apiCategory,
		FunctionName: functionName,
		EmptyValue:   struct{}{},
	}

	if serviceError != nil {
		baseResponse.Error = serviceError
	} else {
		baseResponse.Error = NewServiceError(0, "")
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
	Error        *ServiceError `thrift:",1"`
	Timestamp    BaseTimestamp `thrift:",2"`
	Category     string        `thrift:",3"`
	FunctionName string        `thrift:",4"`
	Data         T             `thrift:",5,omitempty"`
	EmptyValue   Empty         `thrift:",6"`
}

type Empty struct{}
