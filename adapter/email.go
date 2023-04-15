package adapter

import (
	"be-ifid/config"
	"be-ifid/utils"
	"crypto/tls"
	"time"

	gomail "github.com/xhit/go-simple-mail/v2"
)

var server *gomail.SMTPServer

func EmailInit() {
	conf := config.Get()
	server = gomail.NewSMTPClient()
	server.Host = conf.Email.Host
	server.Port = conf.Email.Port
	server.Username = conf.Email.User
	server.Password = conf.Email.Password
	server.Encryption = gomail.EncryptionTLS
	server.KeepAlive = true
	server.ConnectTimeout = 30 * time.Second
	server.SendTimeout = 30 * time.Second
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	utils.Info("Config Email server...")
}

func EmailConnection() *gomail.SMTPServer {
	return server
}
