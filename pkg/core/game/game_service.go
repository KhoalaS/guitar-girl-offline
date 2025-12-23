package game

import (
	"time"
)

type GameService interface {
	Init(params InitParameters) (InitServerUrls, *ServiceError)
	GetServerTime(params GetServerTimeParams) (BaseTimestamp, *ServiceError)
	UserLogin(params UserLoginParams) (*UserLoginResult, *ServiceError)
}

type GameServiceImpl struct{}

func (service *GameServiceImpl) Init(params InitParameters) (InitServerUrls, *ServiceError) {
	return InitServerUrls{
		UnknownInt:     253,
		GameServerUrl:  "https://game.gtgl.pmang.cloud",
		AssetServerUrl: "https://dl.gtgl.pmang.cloud",
	}, nil
}

func (service *GameServiceImpl) GetServerTime(params GetServerTimeParams) (BaseTimestamp, *ServiceError) {
	return getBaseTimeStamp(time.Now(), params.TimeZone), nil
}

func (service *GameServiceImpl) UserLogin(params UserLoginParams) (*UserLoginResult, *ServiceError) {
	// TODO internal logic
	return nil, NewServiceError(
		998, "Sorry. Not a registered user.",
	)
}

type ServiceError struct {
	ErrorCode    int    `thrift:",1"`
	ErrorMessage string `thrift:",2"`
}

func NewServiceError(code int, message string) *ServiceError {
	return &ServiceError{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

type UserLoginResult struct {
	ResultCode    int    `thrift:",1"`
	ResultMessage string `thrift:",2"`
}

type UserLoginParams struct {
	UnknownInt    int32
	UnknownString string
	UserId        string
	DeviceId      string
}

type InitParameters struct {
	Version     int32
	Environment string
	UnknownFlag int16
	DeviceId    string
}

type InitServerUrls struct {
	UnknownInt     int    `thrift:",1"`
	GameServerUrl  string `thrift:",2"`
	AssetServerUrl string `thrift:",3"`
}

type GetServerTimeParams struct {
	DeviceId string
	TimeZone string
}
