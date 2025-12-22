package game

import (
	"github.com/rs/zerolog/log"
	"github.com/thrift-iterator/go/general"
	"github.com/thrift-iterator/go/protocol"
)

func EnvironmentMapperFunc(data general.Struct) EnvironmentData {
	return EnvironmentData{
		Version:     data.Get(protocol.FieldId(1)).(int32),
		Environment: data.Get(protocol.FieldId(2)).(string),
		UnknownFlag: data.Get(protocol.FieldId(3)).(int16),
	}
}

func InitParametersMapperFunc(data general.Struct) InitParameters {
	log.Debug().Any("initparamters data", data).Send()
	return InitParameters{
		Environment: data.Get(protocol.FieldId(1)).(string),
		UnknownFlag: data.Get(protocol.FieldId(2)).(int16),
		Version:     data.Get(protocol.FieldId(3)).(int32),
		DeviceId:    data.Get(protocol.FieldId(4)).(string),
	}
}

func GetServerTimeParamsMapperFunc(data general.Struct) GetServerTimeParams {
	return GetServerTimeParams{
		DeviceId: data.Get(protocol.FieldId(4)).(string),
	}
}
