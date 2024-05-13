package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Favorite struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       string             `json:"user_id" validate:"required"`
	ProductID    string             `json:"product_id" validate:"required"`
	FavoriteDate time.Time          `json:"favorite_date"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
