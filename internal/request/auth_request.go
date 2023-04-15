package request

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuthRequest struct {
	Email string `json:"email"`
}

func (req AuthRequest) Validate() error {
	return validation.ValidateStructWithContext(context.Background(), &req,
		validation.Field(&req.Email, validation.Required, is.EmailFormat),
	)
}
