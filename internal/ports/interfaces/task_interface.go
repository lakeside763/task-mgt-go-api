package interfaces

import "github.com/lakeside763/task-mgt/internal/domain/task"

type TasksRepository interface {
	Create(task *task.Task) (*task.Task, error)
	GetById(id string) (*task.Task, error)
	Update(task *task.Task) (*task.Task, error)
	Delete(id string) error
}