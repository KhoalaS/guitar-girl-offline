package rpc

import (
	"github.com/KhoalaS/thrifter"
	"github.com/KhoalaS/thrifter/general"
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

func ThriftBytesToBz2B64(input []byte) (string, error) {
	compressedBytes, err := compressBz2(input)
	if err != nil {
		return "", err
	}

	return encodeBase64(compressedBytes), nil
}
