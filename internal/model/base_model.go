package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenJWT struct {
	Data interface{} `json:"data"`
	jwt.RegisteredClaims
}
