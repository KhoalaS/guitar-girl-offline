package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/KhoalaS/guitar-girl-offline/pkg/rpc"
	"github.com/KhoalaS/thrifter"
	"github.com/KhoalaS/thrifter/general"
	"github.com/KhoalaS/thrifter/protocol"
)

func main() {

	serverUrl := "https://game.gtgl.pmang.cloud"
	client := http.Client{}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyUrl := serverUrl + r.URL.Path
		r.ParseForm()

		fmt.Println(proxyUrl)

		proxyForm := url.Values{}
		proxyForm.Add("call", r.Form.Get("call"))
		proxyForm.Add("tapsonic_data", r.Form.Get("tapsonic_data"))
		proxyForm.Add("access_token", r.Form.Get("access_token"))
		proxyForm.Add("Cache-Control", r.Form.Get("Cache-Control"))
		proxyForm.Add("current_time", r.Form.Get("current_time"))

		proxyReq, _ := http.NewRequest(r.Method, proxyUrl, strings.NewReader(proxyForm.Encode()))

		for key, header := range r.Header {
			proxyReq.Header.Add(key, strings.Join(header, ", "))
		}

		res, err := client.Do(proxyReq)
		if err != nil {
			w.Write([]byte("error"))
			return
		}

		data, _ := io.ReadAll(res.Body)

		generalStruct, _ := rpc.ThriftDataToStruct(string(data))
		generalStruct[protocol.FieldId(6)] = general.Struct{}

		responseData, _ := thrifter.Marshal(generalStruct)
		encodedData, _ := rpc.ThriftBytesToBz2B64(responseData)

		w.Write([]byte(encodedData))
	})

	server := http.Server{
		Addr:    ":10003",
		Handler: mux,
	}

	server.ListenAndServe()
}
