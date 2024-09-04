package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	// "github.com/lakeside763/task-mgt/internal/infra/cache"
	"github.com/lakeside763/task-mgt/internal/adapters/repository"
	"github.com/lakeside763/task-mgt/internal/ports/http/routes"
	log "github.com/sirupsen/logrus"
)

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main () {
	// Initialize environment variables
	initEnv()

	// Initialize repositories
	repo := repository.NewRepository()

	// redisCache := cache.NewRedisCache(config.RedisURL)

	// Initialize router and setup routes
	router := httprouter.New();
	routes.SetupRoutes(router, repo)

	log.Infof("Starting server on :5200")
	log.Fatal(http.ListenAndServe(":5200", router))
}


