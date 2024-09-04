package mongodb

import (
	"context"
	"time"

	"github.com/lakeside763/task-mgt/internal/domain/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(client *mongo.Client, dbName string) *MongoTaskRepository {
	return &MongoTaskRepository{
		collection: client.Database(dbName).Collection("tasks"),
	}
}

func (r *MongoTaskRepository) Create(task *task.Task) (*task.Task, error) {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		task.ID = oid.Hex()
	} else {
		return nil, mongo.ErrInvalidIndexValue
	}

	return task, nil
}

// Delete implements task.Repository.
func (r *MongoTaskRepository) Delete(id string) error {
	panic("unimplemented")
}

// GetById implements task.Repository.
func (r *MongoTaskRepository) GetById(id string) (*task.Task, error) {
	panic("unimplemented")
}

// Update implements task.Repository.
func (r *MongoTaskRepository) Update(task *task.Task) (*task.Task, error) {
	panic("unimplemented")
}

