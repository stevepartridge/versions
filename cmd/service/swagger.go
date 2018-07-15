package main

import (
	"mime"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/go-openapi/spec"

	"github.com/stevepartridge/versions/swagger"
)

func serveSwagger() {

	service.Router.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, req *http.Request) {

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

	docsPath := "/docs"

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "static/swagger-ui",
	})

	fs := http.StripPrefix(docsPath, fileServer)

	if docsPath != "/" && docsPath[len(docsPath)-1] != '/' {
		service.Router.Get(docsPath, http.RedirectHandler(docsPath+"/", 301).ServeHTTP)
		docsPath += "/"
	}
	docsPath += "*"

	service.Router.Get(docsPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

}

func errorJSON(err error) []byte {
	return []byte(`
{
	"error": "` + err.Error() + `"
}
	`)
}
