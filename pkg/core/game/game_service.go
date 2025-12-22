package game

import "time"

type GameService interface {
	Init(params InitParameters) InitServerUrls
	GetServerTime(params GetServerTimeParams) BaseTimestamp
}

type GameServiceImpl struct{}

func (service *GameServiceImpl) Init(params InitParameters) InitServerUrls {
	return InitServerUrls{
		UnknownInt:     253,
		GameServerUrl:  "https://game.gtgl.pmang.cloud",
		AssetServerUrl: "https://dl.gtgl.pmang.cloud",
	}
}

func (service *GameServiceImpl) GetServerTime(params GetServerTimeParams) BaseTimestamp {
	return getBaseTimeStamp(time.Now(), params.TimeZone)
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
