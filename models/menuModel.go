package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Menu represents a menu entity
type Menu struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name      string               `json:"name" validate:"required"`
	Category  string               `json:"category" validate:"required"`
	StartDate *time.Time           `json:"start_date"`
	EndDate   *time.Time           `json:"end_date"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	MenuID    string               `json:"menu_id"`
	Products  []primitive.ObjectID `bson:"products" json:"products"`
}
