package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/task-mgt/internal/adapters/repository"
)

func SetupRoutes(router *httprouter.Router, repo *repository.Repository) {
	TaskRoutes(router, repo)
}