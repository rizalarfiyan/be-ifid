package repository

import (
	"be-ifid/config"
	"be-ifid/database"
	"context"

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
