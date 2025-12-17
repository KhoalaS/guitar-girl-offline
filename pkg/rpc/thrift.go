package rpc

import (
	thrifter "github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go/general"
)

func ThriftDataToStruct(input string) (general.Struct, error) {
	decocdedBytes, err := decodeBase64(input)
	if err != nil {
		return nil, err
	}

	decompressedBytes, err := decompressBzip2(decocdedBytes)
	if err != nil {
		return nil, err
	}

	var thriftStruct general.Struct
	err = thrifter.Unmarshal(decompressedBytes, &thriftStruct)
	if err != nil {
		return nil, err
	}

	return thriftStruct, nil
}
