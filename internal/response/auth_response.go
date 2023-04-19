package response

import (
	"strings"

	"github.com/google/uuid"
)

type AuthTokenResponse struct {
	Token string `json:"token"`
	IsNew bool   `json:"is_new"`
}

type AuthMeResponse struct {
	ID        *uuid.UUID `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	FullName  string     `json:"full_name"`
	IsValid   bool       `json:"is_new"`
}

func (auth *AuthMeResponse) GetFullName() {
	auth.FullName = auth.FirstName
	if strings.TrimSpace(auth.LastName) != "" {
		auth.FullName += " " + auth.LastName
	}
}
