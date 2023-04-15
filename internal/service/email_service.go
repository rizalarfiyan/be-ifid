package service

import "be-ifid/internal/model"

type EmailService interface {
	SendEmail(payload model.MailPayload, template string, data interface{}) error
}
