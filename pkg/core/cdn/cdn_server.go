package cdn

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed Android.manifest
var androidManifest []byte

//go:embed table_bundlesize.ab.manifest
var tableBundlesizeManifest []byte

func NewCdnServer(port uint16) http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /AssetBundles/Android/Android.manifest", func(w http.ResponseWriter, r *http.Request) {
		w.Write(androidManifest)
	})

	mux.HandleFunc("GET /AssetBundles/Android/table/table_bundlesize.ab.manifest HTTP/1.1", func(w http.ResponseWriter, r *http.Request) {
		w.Write(tableBundlesizeManifest)
	})

	return http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
}
