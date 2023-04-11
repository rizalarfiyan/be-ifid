package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         int
	ServerHost         string
	FirebaseConfig     FirebaseConfig
	FirebaseCredential FirebaseCredentials
}

type FirebaseConfig struct {
	DatabaseUrl string
}

type FirebaseCredentials struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

var (
	conf *Config
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env is not loaded properly")
	}

	conf = new(Config)
	conf.ServerPort, _ = strconv.Atoi(os.Getenv(`PORT`))
	conf.ServerHost = os.Getenv(`HOST`)

	conf.FirebaseConfig.DatabaseUrl = os.Getenv(`FIREBASE_DATABASE`)
	conf.FirebaseCredential.Type = os.Getenv(`FIREBASE_TYPE`)
	conf.FirebaseCredential.ProjectId = os.Getenv(`FIREBASE_PROJECT_ID`)
	conf.FirebaseCredential.PrivateKeyId = os.Getenv(`FIREBASE_PRIVATE_KEY_ID`)
	conf.FirebaseCredential.PrivateKey = os.Getenv(`FIREBASE_PRIVATE_KEY`)
	conf.FirebaseCredential.ClientEmail = os.Getenv(`FIREBASE_CLIENT_EMAIL`)
	conf.FirebaseCredential.ClientId = os.Getenv(`FIREBASE_CLIENT_ID`)
	conf.FirebaseCredential.AuthUri = os.Getenv(`FIREBASE_AUTH_URI`)
	conf.FirebaseCredential.TokenUri = os.Getenv(`FIREBASE_TOKEN_URI`)
	conf.FirebaseCredential.AuthProviderX509CertUrl = os.Getenv(`FIREBASE_AUTH_PROVIDER_X509_CERT_URL`)
	conf.FirebaseCredential.ClientX509CertUrl = os.Getenv(`FIREBASE_CLIENT_X509_CERT_URL`)
}

func Get() *Config {
	return conf
}
