package rpc

import (
	"bytes"
	"compress/bzip2"
	"encoding/base64"
	"io"
)

func decodeBase64(input string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}

	return decodedBytes, nil
}

func decompressBzip2(input []byte) ([]byte, error) {
	inputReader := bytes.NewReader(input)
	decompressedReader := bzip2.NewReader(inputReader)

	decompressedBuffer, err := io.ReadAll(decompressedReader)
	if err != nil {
		return nil, err
	}

	return decompressedBuffer, nil
}
