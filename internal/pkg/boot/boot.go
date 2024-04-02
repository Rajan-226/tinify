package boot

import (
	"context"
	"github.com/myProjects/tinify/internal/app/controllers"
	"github.com/myProjects/tinify/internal/app/tinify"
	"github.com/myProjects/tinify/models/url_info"
	redis "github.com/redis/go-redis/v9"
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

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	tinify.NewCore(urlCore, client)
}
