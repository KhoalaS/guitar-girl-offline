package rpc

import (
	"os"
	"testing"

	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
	"github.com/KhoalaS/thrifter"
	"github.com/KhoalaS/thrifter/protocol"
)

func TestThriftDataToStruct(t *testing.T) {
	testee, _ := os.ReadFile("./test_data/raw_data.bz2.b64")

	actual, err := ThriftDataToStruct(string(testee))
	if err != nil {
		t.Error(err)
	}

	if actual == nil {
		t.Fail()
	}

	action := actual.Get(protocol.FieldId(1))
	if action != "init" {
		t.Fail()
	}

	productionEnv := actual.Get(protocol.FieldId(2), protocol.FieldId(1))
	if productionEnv != "real" {
		t.Fail()
	}

	flag := actual.Get(protocol.FieldId(2), protocol.FieldId(2))
	if flag != int16(1) {
		t.Fail()
	}

	version := actual.Get(protocol.FieldId(2), protocol.FieldId(3))
	if version != int32(800) {
		t.Fail()
	}
}

func TestThriftMapEncoding(t *testing.T) {
	m := map[string]string{}
	m["ret"] = "1"

	_, err := thrifter.Marshal(&m)
	if err != nil {
		t.Error(err)
	}
}

func TestThriftMarshalling(t *testing.T) {
	message := "QlpoMzFBWSZTWaNUnfgAABv9oP9uAQBAAH/gIAA/TRwAAAIAQCAAdCKb1J5TJkJoxMmCaNp6hntIJiAAAGgNDUk0qCiAJrFg4nBpqQACeCjBEXLHIiohjds9ZQeP0+xZaU/Mk0nYZCotc+G7pBIMDqImEOlD3MLV/uNoQPlw0RxPpjuorY4lKEsRQua9i7kinChIUapO/AA="
	requestData, err := ThriftDataToAny[user_model.SetAttendance](message)
	if err != nil {
		t.Error(err)
	}

	if requestData.Call != "setAttendance" {
		t.Fail()
	}
}
