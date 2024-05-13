package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name              string             `json:"name" validate:"required,min=2,max=100"`
	Description       string             `json:"description" validate:"required,min=2,max=600"`
	Roasted           string             `json:"roasted" validate:"required,min=2,max=100"`
	ImageLinkSquare   string             `json:"image_link_square" validate:"required"`
	ImageLinkPortrait string             `json:"image_link_portrait" validate:"required"`
	Ingredients       string             `json:"ingredients"`
	SpecialIngredient string             `json:"special_ingredient"`
	Prices            []Price            `json:"prices" validate:"required"`
	AverageRating     *float64           `json:"average_rating"`
	RatingsCount      *string            `json:"ratings_count"`
	Favourite         *bool              `json:"favourite"`
	Type              *string            `json:"type"`
	Index             *int               `json:"index"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	FoodID            string             `json:"food_id"`
	MenuID            string             `json:"menu_id" validate:"required"`
}
