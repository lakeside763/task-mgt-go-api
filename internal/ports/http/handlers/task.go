package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/task-mgt/internal/domain/task"
	"github.com/lakeside763/task-mgt/internal/ports/interfaces"
	"github.com/lakeside763/task-mgt/pkg/utils"
)

type TaskHandler struct {
	TaskRepo interfaces.TasksRepository
}

func NewTaskHandler(repo interfaces.TasksRepository) *TaskHandler {
	return &TaskHandler{TaskRepo: repo}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var task task.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newTask, err := h.TaskRepo.Create(&task); 
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, http.StatusOK, newTask)
}