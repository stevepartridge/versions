package main

import (
	"mime"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/go-openapi/spec"

	"github.com/stevepartridge/versions/swagger"
)

func serveSwagger(mux *http.ServeMux) {

	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, req *http.Request) {

		data, err := swagger.Asset("static/swagger-ui/service.swagger.json")
		if err != nil {
			ifError(err)
			w.Write(errorJSON(err))
		}

		swag := spec.Swagger{}

		err = swag.UnmarshalJSON(data)
		if err != nil {
			ifError(err)
			w.Write(errorJSON(err))
		}

		swag.Info.Title = serviceName
		swag.Info.Version = version
		swag.Schemes = []string{"https"}

		data, err = swag.MarshalJSON()
		if err != nil {
			ifError(err)
			w.Write(errorJSON(err))
		}

		w.Write(data)
	})

	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in static/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "static/swagger-ui",
	})

	prefix := "/docs/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func errorJSON(err error) []byte {
	return []byte(`
{
	"error": "` + err.Error() + `"
}
	`)
}
