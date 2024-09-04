package postgresql

import (
	"github.com/lakeside763/task-mgt/internal/domain/task"
	"gorm.io/gorm"
)

type PostgresTaskRepository struct {
	db *gorm.DB
}

func NewPostgresTaskRepository(db *gorm.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(task *task.Task) (*task.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete implements task.Repository.
func (r *PostgresTaskRepository) Delete(id string) error {
	panic("unimplemented")
}

// GetById implements task.Repository.
func (r *PostgresTaskRepository) GetById(id string) (*task.Task, error) {
	panic("unimplemented")
}

// Update implements task.Repository.
func (r *PostgresTaskRepository) Update(task *task.Task) (*task.Task, error) {
	panic("unimplemented")
}


