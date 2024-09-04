package main

import (
	"context"
	"errors"

	"github.com/lakeside763/task-mgt/config"
	"github.com/lakeside763/task-mgt/internal/adapters/cache"
	"github.com/lakeside763/task-mgt/internal/adapters/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabaseAndCache(config config.Config) (*gorm.DB, *cache.RedisCache, error) {
	var db *gorm.DB
	var err error

	switch config.DatabaseType {
	case "mongodb":
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DatabaseURL))
		if err != nil {
			return nil, nil, err
		}
		_ = mongodb.NewMongoTaskRepository(client, "task_mgt2")
	case "postgresql":
		db, err = gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("unsupported database type")
	}

	redisCache := cache.NewRedisCache(config.RedisURL)
	return db, redisCache, nil
}