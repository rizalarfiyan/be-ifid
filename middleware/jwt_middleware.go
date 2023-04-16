package middleware

import (
	"be-ifid/config"
	"be-ifid/internal/response"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	Decode     func(c *fiber.Ctx) (*jwt.MapClaims, error)
	LocalToken func(ctx *fiber.Ctx, claims jwt.MapClaims) error
	Secret     string
}

var JWTConfigDefault = JWTConfig{
	Decode: nil,
}

func configJWTDefault(conf ...JWTConfig) JWTConfig {
	if len(conf) < 1 {
		return JWTConfigDefault
	}

	cfg := conf[0]
	if cfg.Secret == "" {
		cfg.Secret = config.Get().JWT.SecretKey
	}

	if cfg.Decode == nil {
		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {
			authHeader := c.Get("Authorization")
			if authHeader == "" {
				return nil, response.NewErrorMessage(http.StatusUnauthorized, "Missing auth header", nil)
			}

			if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
				return nil, response.NewErrorMessage(http.StatusUnauthorized, "Invalid auth header", nil)
			}

			token, err := jwt.Parse(
				authHeader[7:],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, response.NewErrorMessage(http.StatusUnauthorized, fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]), nil)
					}
					return []byte(cfg.Secret), nil
				},
			)
			if err != nil {
				return nil, response.NewErrorMessage(http.StatusUnauthorized, "Error parsing token", nil)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !(ok && token.Valid) {
				return nil, response.NewErrorMessage(http.StatusUnauthorized, "Invalid token", nil)
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				return nil, response.NewErrorMessage(http.StatusUnauthorized, "JWT is expired", nil)
			}

			return &claims, nil
		}
	}

	if cfg.LocalToken == nil {
		cfg.LocalToken = func(ctx *fiber.Ctx, claims jwt.MapClaims) error {
			claimData := claims["data"]
			if claimData == nil {
				return response.NewErrorMessage(http.StatusUnprocessableEntity, "JWT payload not have data", nil)
			}

			ctx.Locals("user", claimData)
			return nil
		}
	}

	return cfg
}

func NewJWTMiddleware(config JWTConfig) fiber.Handler {
	cfg := configJWTDefault(config)

	return func(ctx *fiber.Ctx) error {
		claims, err := cfg.Decode(ctx)
		if err != nil {
			return err
		}

		err = cfg.LocalToken(ctx, *claims)
		if err != nil {
			return err
		}

		return ctx.Next()
	}
}
