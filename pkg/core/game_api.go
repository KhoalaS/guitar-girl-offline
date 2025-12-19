package core

import "time"

type GameApi struct {
	service GameService
}

func NewGameApi() GameApi {
	return GameApi{
		service: &GameServiceImpl{},
	}
}

func (api *GameApi) Init(params BaseGameRequest[InitParameters], apiCategory string) BaseGameResponse[InitServerUrls] {
	serverList := api.service.Init(params.Data)
	return NewBaseGameResponse(params.FunctionName, apiCategory, serverList)
}

func NewBaseGameResponse[T any](functionName string, apiCategory string, data T) BaseGameResponse[T] {
	now := time.Now()
	return BaseGameResponse[T]{
		UnknownField: UnknownField{
			UnknownInt:    0,
			UnknownString: "",
		},
		Timestamp: BaseTimestamp{
			UnixSeconds:       now.Unix(),
			ServerTimeIsoDate: int64(MustAtoi(SecondsToServerISO(now))),
		},
		Category:     apiCategory,
		FunctionName: functionName,
		EmptyValue:   struct{}{},
		Data:         data,
	}
}

func SecondsToServerISO(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Seoul")
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
	Version     int
	Environment string
	UnknownFlag int
}

type BaseGameRequest[T any] struct {
	FunctionName string
	Data         T
	Environment  EnvironmentData
}

type BaseGameResponse[T any] struct {
	UnknownField UnknownField
	Timestamp    BaseTimestamp
	Category     string
	FunctionName string
	EmptyValue   any
	Data         T
}
