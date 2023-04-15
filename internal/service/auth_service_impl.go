package service

import (
	"be-ifid/adapter"
	"be-ifid/config"
	"be-ifid/constant"
	"be-ifid/database"
	"be-ifid/internal/model"
	"be-ifid/internal/repository"
	"be-ifid/internal/request"
	"be-ifid/internal/response"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (s *authService) createToken(data model.JWTAuthPayload) (string, error) {
	claims := model.TokenJWT{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.conf.JWT.Expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.conf.JWT.SecretKey))
}

func (s *authService) Callback(token string) (*response.AuthCallbackResponse, error) {
	keySearch := fmt.Sprintf("%s*:%s", constant.RedisKeyAuth, token)
	searchKeys, err := s.redis.Keys(keySearch)
	if err != nil {
		return nil, err
	}

	if len(searchKeys) != 1 {
		return nil, response.NewErrorMessage(http.StatusUnprocessableEntity, "Token expired!", nil)
	}

	email, err := s.redis.GetString(searchKeys[0])
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	payload := model.JWTAuthPayload{
		Email: email,
		IsNew: true,
	}

	if user != nil {
		payload.ID = &user.ID
		payload.Email = user.Email
		payload.FirstName = user.FirstName
		payload.LastName = user.LastName
		payload.IsNew = false
	}

	jwtToken, err := s.createToken(payload)
	if err != nil {
		return nil, err
	}

	return &response.AuthCallbackResponse{
		IsNew: payload.IsNew,
		Token: jwtToken,
	}, nil
}
