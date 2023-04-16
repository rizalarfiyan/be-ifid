package request

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginRequest struct {
	Email string `json:"email"`
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStructWithContext(context.Background(), &req,
		validation.Field(&req.Email, validation.Required, is.EmailFormat),
	)
}

type FirstUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (req FirstUserRequest) Validate() error {
	return validation.ValidateStructWithContext(context.Background(), &req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.LastName),
	)
}
