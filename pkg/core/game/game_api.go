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
	serverList := api.service.Init(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, api.timeZone, serverList)
}

func NewBaseGameResponse[T any](functionName string, apiCategory string, timeZone string, data T) BaseGameResponse[T] {
	now := time.Now()
	return BaseGameResponse[T]{
		UnknownField: UnknownField{
			UnknownInt:    0,
			UnknownString: "",
		},
		Timestamp: BaseTimestamp{
			UnixSeconds:       now.Unix(),
			ServerTimeIsoDate: int64(core.MustAtoi(SecondsToServerISO(now, timeZone))),
		},
		Category:     apiCategory,
		FunctionName: functionName,
		EmptyValue:   struct{}{},
		Data:         data,
	}
}

func SecondsToServerISO(t time.Time, timeZone string) string {
	loc, _ := time.LoadLocation(timeZone)
	return t.In(loc).Format("20060102150405")
}

type UnknownField struct {
	UnknownInt    int
	UnknownString string
}

type BaseTimestamp struct {
	UnixSeconds       int64
	ServerTimeIsoDate int64
}

type EnvironmentData struct {
	Version     int32
	Environment string
	UnknownFlag int16
}

type BaseGameRequest[T any] struct {
	FunctionName string
	Data         T
	Environment  EnvironmentData
}

type BaseGameResponse[T any] struct {
	UnknownField UnknownField  `thrift:",1"`
	Timestamp    BaseTimestamp `thrift:",2"`
	Category     string        `thrift:",3"`
	FunctionName string        `thrift:",4"`
	Data         T             `thrift:",5"`
	EmptyValue   Empty         `thrift:",6"`
}

type Empty struct{}
