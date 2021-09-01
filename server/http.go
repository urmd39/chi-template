package server

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"

	"net/http"
	_ "nutrition/docs"
	"nutrition/infrastructure"
	"nutrition/server/api"
	"strings"
)

func Router() http.Handler {
	r := chi.NewRouter()
	basePath := "/api/v1/receivable"
	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.HttpSwagger), //The url pointing to API definition"
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Homepage"))
	})

	r.Route(basePath, func(subR chi.Router) {
		subR.Route("", api.NewExampleRouter)
	})
	return r
}

func HandleHttpServer(port string) {
	infrastructure.InfoLog.Printf("server  %s", port)
	infrastructure.InfoLog.Printf("swagger : %s", strings.Replace(infrastructure.HttpSwagger, "doc.json", "index.html", 1))
	http.ListenAndServe(port, Router())
}
