package service

import (
	"be-ifid/adapter"
	"be-ifid/config"
	"be-ifid/constant"
	"be-ifid/database"
	"be-ifid/internal/model"
	"be-ifid/internal/repository"
	"be-ifid/internal/request"
	"context"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type authService struct {
	ctx   context.Context
	conf  *config.Config
	repo  repository.AuthRepository
	redis database.RedisInstance
}

func NewAuthService(ctx context.Context, conf *config.Config, repo repository.AuthRepository, redis database.RedisInstance) AuthService {
	return &authService{
		ctx,
		conf,
		repo,
		redis,
	}
}

func (s *authService) getUniqueKey() string {
	id, _ := gonanoid.New(constant.AuthKeyLength)
	return id
}

func (s *authService) sendVerificationEmail(email string, key string) error {
	emailConnection := adapter.EmailConnection()
	payload := model.MailPayload{
		From:    s.conf.Email.From,
		To:      email,
		Subject: "Welcome To IFID!",
	}

	var data = make(map[string]interface{})
	data["fullName"] = "Rizal Arfiyan"
	data["verificationCode"] = s.conf.FE.BaseUrl + s.conf.FE.AuthRedirectUrl + "?token=" + key
	return NewEmailService(s.conf, emailConnection).SendEmail(payload, constant.TemplateSignup, data)
}

func (s *authService) Login(req request.AuthRequest) error {
	keyUnique := s.getUniqueKey()
	keyRedis := fmt.Sprintf("%s%s:%s", constant.RedisKeyAuth, req.Email, keyUnique)
	err := s.redis.Setxc(keyRedis, constant.AuthExpire, req.Email)
	if err != nil {
		return err
	}

	err = s.sendVerificationEmail(req.Email, keyUnique)
	if err != nil {
		return err
	}

	return nil
}
