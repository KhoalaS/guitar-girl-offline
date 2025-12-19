package core

type GameService interface {
	Init(params InitParameters) InitServerUrls
}

type GameServiceImpl struct{}

func (service *GameServiceImpl) Init(params InitParameters) InitServerUrls {
	return InitServerUrls{
		UnknownInt:     253,
		GameServerUrl:  "https://game.gtgl.pmang.cloud",
		AssetServerUrl: "https://dl.gtgl.pmang.cloud",
	}
}

type InitParameters struct {
	Version     int
	Environment string
	UnknownFlag int
	DeviceId    string
}

type InitServerUrls struct {
	UnknownInt     int
	GameServerUrl  string
	AssetServerUrl string
}
