package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/task-mgt/internal/adapters/repository"
	"github.com/lakeside763/task-mgt/internal/ports/http/handlers"
)

func TaskRoutes(router *httprouter.Router, repo *repository.Repository) {
	taskHandler := handlers.NewTaskHandler(repo.TaskRepo)

	router.POST("/tasks", taskHandler.CreateTask)
}