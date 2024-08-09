package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID primitive.ObjectID 	`json:"_id" validate:"required,min=2"`
	Name *string `json:"name" validate:"required,min=2,max=25"` 
	Password *string `json:"password" validate:"required min=6,max=100"`
	Email *string `json:"email" validate:"required email"`
	Token *string `json:"token"`
	Role *string `json:"role"`
	RefreshToken *string `json:"refresh_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
	UserID string `json:"user_id"`
}


type PromoteUserRequest struct {
    ID string `json:"id" binding:"required"`
}
