package service

import "be-ifid/internal/model"

type EmailService interface {
	SendEmail(payload model.MailPayload) error
}
