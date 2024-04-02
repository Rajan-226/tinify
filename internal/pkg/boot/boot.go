package boot

import (
	"context"
	"github.com/myProjects/tinify/internal/app/controllers"
	"github.com/myProjects/tinify/internal/app/tinify"
	"github.com/myProjects/tinify/models/domain_info"
	"github.com/myProjects/tinify/models/url_info"
	"net/http"

	gmux "github.com/gorilla/mux"
)

func Init(ctx context.Context) {
	server := NewServer(ctx)

	initEntities()

	server.ListenAndServe()

}

func NewServer(ctx context.Context) *http.Server {
	mux := gmux.NewRouter()

	//middlewares

	mux.Methods(http.MethodPost).Path("/v1/tinify").HandlerFunc(controllers.Tinify)
	mux.Methods(http.MethodGet).Path("/v1/metrics").HandlerFunc(controllers.Metrics)
	mux.Methods(http.MethodGet).Path("/v1/{path}").HandlerFunc(controllers.Redirect)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	return server
}

func initEntities() {
	urlCore := url_info.NewCore(url_info.NewRepo())
	domainCore := domain_info.NewCore(domain_info.NewRepo())

	tinify.NewCore(urlCore, domainCore)
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
