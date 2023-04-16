package model

import (
	"be-ifid/config"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthIdentity struct {
	ID               *uuid.UUID `json:"id"`
	Email            string     `json:"email"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	FullName         string     `json:"full_name"`
	IsNew            bool       `json:"is_new"`
	VerificationCode string     `json:"verification_code"`
}

type JWTAuthPayload struct {
	ID        *uuid.UUID `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	IsNew     bool       `json:"is_new"`
}

func (auth *AuthIdentity) GetFullName() {
	auth.FullName = fmt.Sprintf("%s %s", auth.FirstName, auth.LastName)
}

func (auth *AuthIdentity) SetVerificationCode(keyUnique string, conf config.Config) {
	auth.VerificationCode = conf.FE.BaseUrl + conf.FE.AuthRedirectUrl + "?token=" + keyUnique
}

func (jwt *JWTAuthPayload) GetFromFiber(ctx *fiber.Ctx) error {
	userStr, err := json.Marshal(ctx.Locals("user"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(userStr, &jwt)
	if err != nil {
		return err
	}
	return nil
}
