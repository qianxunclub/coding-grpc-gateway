package swagger

import (
	"coding-grpc-gateway/pkg/ui/data/swagger"
	"github.com/elazarl/go-bindata-assetfs"
	"log"
	"net/http"
	"path"
	"strings"
)

func ServeSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("swagger/json", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func ServeSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
