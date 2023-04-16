package service

import (
	"be-ifid/config"
	"be-ifid/internal/model"
	"bytes"
	"text/template"

	gomail "github.com/xhit/go-simple-mail/v2"
)

type emailService struct {
	conf   *config.Config
	server *gomail.SMTPServer
}

func NewEmailService(conf *config.Config, email *gomail.SMTPServer) EmailService {
	return &emailService{
		conf,
		email,
	}
}

func (s *emailService) generateTemplate(templatePath string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func (s *emailService) SendEmail(payload model.MailPayload) error {
	template, err := s.generateTemplate(payload.Template, payload.Data)
	if err != nil {
		return err
	}

	smtpClient, err := s.server.Connect()
	if err != nil {
		return err
	}

	email := gomail.NewMSG()
	email.SetFrom(payload.From).AddTo(payload.To).SetSubject(payload.Subject)
	email.SetBody(gomail.TextHTML, template)
	if email.Error != nil {
		return err
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}
