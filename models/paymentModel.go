package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PaymentID      string             `json:"invoice_id"`
	OrderID        string             `json:"order_id"`
	PaymentMethod  string             `json:"payment_method" validate:"omitempty,eq=CARD|eq=CASH"`
	PaymentStatus  string             `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time          `json:"payment_due_date"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
