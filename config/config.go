package config

import (
	"be-ifid/helpers"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         int
	ServerHost         string
	FirebaseConfig     FirebaseConfigs
	FirebaseCredential FirebaseCredentials
	Cors               CorsConfigs
	MQTT               MQTTConfigs
}

type FirebaseConfigs struct {
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

type CorsConfigs struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
	ExposeHeaders    string
}

type MQTTConfigs struct {
	ClientId string
	Server   string
	Port     int
	User     string
	Password string
	Topic    struct {
		Watch string
		Send  string
	}
}

var conf *Config

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env is not loaded properly")
	}

	conf = new(Config)
	conf.ServerPort = helpers.GetEnvAsInt("PORT", 8080)
	conf.ServerHost = helpers.GetEnv("HOST", "")

	conf.FirebaseConfig.DatabaseUrl = helpers.GetEnv("FIREBASE_DATABASE", "")
	conf.FirebaseCredential.Type = helpers.GetEnv("FIREBASE_TYPE", "")
	conf.FirebaseCredential.ProjectId = helpers.GetEnv("FIREBASE_PROJECT_ID", "")
	conf.FirebaseCredential.PrivateKeyId = helpers.GetEnv("FIREBASE_PRIVATE_KEY_ID", "")
	conf.FirebaseCredential.PrivateKey = helpers.GetEnv("FIREBASE_PRIVATE_KEY", "")
	conf.FirebaseCredential.ClientEmail = helpers.GetEnv("FIREBASE_CLIENT_EMAIL", "")
	conf.FirebaseCredential.ClientId = helpers.GetEnv("FIREBASE_CLIENT_ID", "")
	conf.FirebaseCredential.AuthUri = helpers.GetEnv("FIREBASE_AUTH_URI", "")
	conf.FirebaseCredential.TokenUri = helpers.GetEnv("FIREBASE_TOKEN_URI", "")
	conf.FirebaseCredential.AuthProviderX509CertUrl = helpers.GetEnv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL", "")
	conf.FirebaseCredential.ClientX509CertUrl = helpers.GetEnv("FIREBASE_CLIENT_X509_CERT_URL", "")

	conf.Cors.AllowOrigins = helpers.GetEnv("ALLOW_ORIGINS", "")
	conf.Cors.AllowMethods = helpers.GetEnv("ALLOW_METHODS", "")
	conf.Cors.AllowHeaders = helpers.GetEnv("ALLOW_HEADERS", "")
	conf.Cors.AllowCredentials = helpers.GetEnvAsBool("ALLOW_CREDENTIALS", false)
	conf.Cors.ExposeHeaders = helpers.GetEnv("EXPOSE_HEADERS", "")

	conf.MQTT.ClientId = helpers.GetEnv("MQTT_CLIENT_ID", "be-ifid")
	conf.MQTT.Server = helpers.GetEnv("MQTT_SERVER", "")
	conf.MQTT.Port = helpers.GetEnvAsInt("MQTT_PORT", 1883)
	conf.MQTT.User = helpers.GetEnv("MQTT_USER", "")
	conf.MQTT.Password = helpers.GetEnv("MQTT_PASSWORD", "")
	conf.MQTT.Topic.Watch = helpers.GetEnv("MQTT_TOPIC_WATCH", "")
	conf.MQTT.Topic.Send = helpers.GetEnv("MQTT_TOPIC_SEND", "")
}

func Get() *Config {
	return conf
}
