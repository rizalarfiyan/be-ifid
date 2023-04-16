package constant

import "time"

const (
	AuthKeyLength = 128
	AuthExpire    = 10 * time.Minute

	RedisKeyAuth = "auth:"

	TemplateSignup = "template/signup.html"
	TemplateLogin  = "template/login.html"
)
