package repository

import (
	"be-ifid/config"
	"be-ifid/database"
	"be-ifid/internal/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
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
	getUserByEmail   = `SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE email = $1`
	countUserByEmail = `SELECT count(*) FROM users WHERE email = $1`
	createUser       = `INSERT INTO users (email, first_name, last_name) VALUES (:email, :first_name, :last_name) RETURNING id`
)

func (r *authRepository) GetUserByEmail(email string) (*model.UserModel, error) {
	user := &model.UserModel{}
	err := r.postgres.GetContext(r.ctx, user, getUserByEmail, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *authRepository) CheckUserByEmail(email string) (bool, error) {
	var count int
	err := r.postgres.GetContext(r.ctx, &count, countUserByEmail, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *authRepository) CreateUser(payload model.CreateUserModel) (uuid.UUID, error) {
	stmt, err := r.postgres.PrepareNamed(createUser)
	if err != nil {
		return uuid.Nil, err
	}

	var id string
	err = stmt.Get(&id, payload)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(id), nil
}
