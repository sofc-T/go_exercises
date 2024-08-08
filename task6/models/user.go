package models 

import (
	"time"
)

type User struct{
	ID *int 	`json:"_id" validate:"required,min=2"`
	Name *string `json:"name" validate:"requried,min=2,max=25"` 
	Password *string `json:"password" validate:"required min=6,max=100"`
	Email *string `json:"email" validate:"requried email"`
	Token *string `json:"token"`
	Role *string `json:"role"`
	RefreshToekn *string `json:"refresh_token"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
	UserID *string `json:"user_id"`
}
