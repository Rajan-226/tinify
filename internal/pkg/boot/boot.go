package boot

import (
	"context"
	"fmt"
	gmux "github.com/gorilla/mux"
	"github.com/myProjects/tinify/internal/app/controllers"
	"github.com/myProjects/tinify/internal/app/tinify"
	"github.com/myProjects/tinify/models/url_info"
	redis "github.com/redis/go-redis/v9"
	"net/http"
)

func Init(ctx context.Context) {
	redisClient := initRedis(ctx)
	initEntities(redisClient)
	initServer(ctx)
}

func initServer(ctx context.Context) {
	mux := gmux.NewRouter()

	//middlewares

	mux.Methods(http.MethodPost).Path("/v1/tinify").HandlerFunc(controllers.Tinify)
	mux.Methods(http.MethodGet).Path("/v1/metrics").HandlerFunc(controllers.Metrics)
	mux.Methods(http.MethodGet).Path("/v1/{path}").HandlerFunc(controllers.Redirect)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	server.ListenAndServe()
}

func initEntities(redisClient redis.UniversalClient) {
	urlCore := url_info.NewCore(url_info.NewRepo())

	tinify.NewCore(urlCore, redisClient)
}

func initRedis(ctx context.Context) redis.UniversalClient {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   2,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("error while pinging redis : %w", err))
	}

	return client
}
