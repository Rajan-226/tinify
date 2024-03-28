package boot

import (
	"context"
	"github.com/myProjects/tinify/internal/app/controllers"
	"net/http"

	gmux "github.com/gorilla/mux"
)

func Init(ctx context.Context) {
	server := NewServer(ctx)

	server.ListenAndServe()

}

func NewServer(ctx context.Context) *http.Server {
	mux := gmux.NewRouter()

	//middlewares

	mux.Methods(http.MethodGet).Path("/v1/tinify").HandlerFunc(controllers.Tinify)
	mux.Methods(http.MethodGet).Path("/v1/redirect").HandlerFunc(controllers.Redirect)
	mux.Methods(http.MethodGet).Path("/v1/metrics").HandlerFunc(controllers.Metrics)

	server := &http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8080",
	}

	return server
}

/*
tinify url API (strategy, url_info db core)
	-> url shortening logic
	-> db entity creation
redirection API (db core)
	-> fetch from core
metrics API ()



->handlers/
	tinify_api.go -> tinify_processor.go


/processors

controllers
	tinify.go
	redirection.go
	metrics.go

api/
	tinify
		GetServer()
		NewServer(NewProcessor(NewCore()))
		server.go -> validations


		processor.go ->

	redirection


	metrics




*/
