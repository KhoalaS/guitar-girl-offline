package core

type GameApi interface {
	Init(data InitRequestData)
}

type InitResponse struct {
	BaseGameResponse
	ServerUrls InitServerUrls
}

type InitRequestData struct {
	EnvironmentData
	DeviceId string
}

type EnvironmentData struct {
	Version     int
	Environment string
	UnknownFlag int
}

type BaseResponseUnknownField struct {
	UnknownInt    int
	UnknownString string
}

type BaseResponseTimestamp struct {
	UnixMillis int64
	IsoDate    string
}

type InitServerUrls struct {
	UnknownInt     int
	GameServerUrl  string
	AssetServerUrl string
}

type BaseGameResponse struct {
	UnknownField BaseResponseUnknownField
	Timestamp    BaseResponseTimestamp
	Type         string
	FunctionName string
	EmptyValue   any
}
