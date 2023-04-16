package repository

import (
	"be-ifid/internal/model"

	"github.com/google/uuid"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*model.UserModel, error)
	CheckUserByEmail(email string) (bool, error)
	CreateUser(payload model.CreateUserModel) (uuid.UUID, error)
}
