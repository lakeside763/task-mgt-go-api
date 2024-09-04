package task

import "time"

type Task struct {
	ID          string    `gorm:"type:uuid;column:id;default:uuid_generate_v4();primaryKey" bson:"_id,omitempty"` // UUID for Postgres, ObjectID for Mongo
	Title       string    `gorm:"size:255;column:title" bson:"title"`
	Description string    `gorm:"type:text;column:description" bson:"description"`
	Status      string    `gorm:"size:255;column:status" bson:"status"`
	Priority    string    `gorm:"size:255;column:priority" bson:"priority"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at" bson:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at" bson:"updatedAt"`
}
