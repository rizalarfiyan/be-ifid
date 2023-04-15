package repository

import "be-ifid/internal/model"

type AuthRepository interface {
	GetUserByEmail(email string) (*model.UserModel, error)
}
