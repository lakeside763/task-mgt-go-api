package repository

import (
	"context"
	"log"

	"github.com/lakeside763/task-mgt/config"
	"github.com/lakeside763/task-mgt/internal/adapters/database/mongodb"
	"github.com/lakeside763/task-mgt/internal/adapters/database/postgresql"
	"github.com/lakeside763/task-mgt/internal/ports/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	TaskRepo interfaces.TasksRepository
}

func NewRepository() *Repository {
	config := config.LoadConfig()

	var taskRepo interfaces.TasksRepository

	// initialize database and manage repository
	switch config.DatabaseType {
	case "mongodb":
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DatabaseURL))
		if err != nil {
			log.Fatal(err)
		}
		taskRepo = mongodb.NewMongoTaskRepository(client, "task_management")
	case "postgresql":
		db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		taskRepo = postgresql.NewPostgresTaskRepository(db)
	default:
		log.Fatal("Unsupported database type")
	}

	return &Repository{
		TaskRepo: taskRepo,
	}
}
