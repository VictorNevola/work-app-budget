package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	BaseModel struct {
		ID        primitive.ObjectID `bson:"_id"`
		CreatedAt time.Time          `bson:"created_at"`
		UpdatedAt time.Time          `bson:"updated_at"`
	}
)

func (b *BaseModel) NewID() {
	id := primitive.NewObjectID()

	b.ID = id
}

func (b *BaseModel) SetCreatedAt() {
	timeNow := time.Now()
	b.CreatedAt = timeNow
	b.UpdatedAt = timeNow
}
