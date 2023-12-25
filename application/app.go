package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	config *Config
	router http.Handler
	redis  *redis.Client
}

func Init() *App {
	config := &Config{}

	app := &App{
		config: config,
		router: Router(),
		redis: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
			DB:   0, // use default DB
		}),
	}

	return app
}

func (app *App) Run(ctx context.Context) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.ServerPort),
		Handler: app.router,
	}

	redisErr := app.redis.Ping(ctx).Err()

	if redisErr != nil {
		fmt.Errorf("Error connecting to redis: %w", redisErr)
	}

	log.Println("connected to redis")
	log.Println("server started at localhost:5000")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Errorf("Error starting server: %w", err)
	}
}
