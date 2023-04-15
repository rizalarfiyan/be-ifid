package repository

import (
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/internal/model"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type authRepository struct {
	ctx      context.Context
	conf     *config.Config
	postgres *sqlx.DB
	redis    database.RedisInstance
}

func NewAuthRepository(ctx context.Context, conf *config.Config, postgres *sqlx.DB, redis database.RedisInstance) AuthRepository {
	return &authRepository{
		ctx,
		conf,
		postgres,
		redis,
	}
}

var (
	getUserByEmail = `SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE email = $1`
)

func (r *authRepository) GetUserByEmail(email string) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := r.postgres.Get(user, getUserByEmail, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
