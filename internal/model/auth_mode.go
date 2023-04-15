package model

import "github.com/google/uuid"

type JWTAuthPayload struct {
	ID        *uuid.UUID `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	IsNew     bool       `json:"is_new"`
}
