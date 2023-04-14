package config

import (
	"be-ifid/helpers"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	Host string
	Cors CorsConfigs
	MQTT MQTTConfigs
	DB   DBConfigs
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

type DBConfigs struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

var conf *Config

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env is not loaded properly")
	}

	conf = new(Config)
	conf.Port = helpers.GetEnvAsInt("PORT", 8080)
	conf.Host = helpers.GetEnv("HOST", "")

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

	conf.DB.Name = helpers.GetEnv("DB_NAME", "")
	conf.DB.Host = helpers.GetEnv("DB_HOST", "")
	conf.DB.Port = helpers.GetEnvAsInt("DB_PORT", 5432)
	conf.DB.User = helpers.GetEnv("DB_USER", "")
	conf.DB.Password = helpers.GetEnv("DB_PASSWORD", "")
}

func Get() *Config {
	return conf
}
