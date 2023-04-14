package config

import (
	"be-ifid/utils"

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
		utils.Error(".env is not loaded properly")
	}

	conf = new(Config)
	conf.Port = utils.GetEnvAsInt("PORT", 8080)
	conf.Host = utils.GetEnv("HOST", "")

	conf.Cors.AllowOrigins = utils.GetEnv("ALLOW_ORIGINS", "")
	conf.Cors.AllowMethods = utils.GetEnv("ALLOW_METHODS", "")
	conf.Cors.AllowHeaders = utils.GetEnv("ALLOW_HEADERS", "")
	conf.Cors.AllowCredentials = utils.GetEnvAsBool("ALLOW_CREDENTIALS", false)
	conf.Cors.ExposeHeaders = utils.GetEnv("EXPOSE_HEADERS", "")

	conf.MQTT.ClientId = utils.GetEnv("MQTT_CLIENT_ID", "be-ifid")
	conf.MQTT.Server = utils.GetEnv("MQTT_SERVER", "")
	conf.MQTT.Port = utils.GetEnvAsInt("MQTT_PORT", 1883)
	conf.MQTT.User = utils.GetEnv("MQTT_USER", "")
	conf.MQTT.Password = utils.GetEnv("MQTT_PASSWORD", "")
	conf.MQTT.Topic.Watch = utils.GetEnv("MQTT_TOPIC_WATCH", "")
	conf.MQTT.Topic.Send = utils.GetEnv("MQTT_TOPIC_SEND", "")

	conf.DB.Name = utils.GetEnv("DB_NAME", "")
	conf.DB.Host = utils.GetEnv("DB_HOST", "")
	conf.DB.Port = utils.GetEnvAsInt("DB_PORT", 5432)
	conf.DB.User = utils.GetEnv("DB_USER", "")
	conf.DB.Password = utils.GetEnv("DB_PASSWORD", "")

	utils.Success("Config is loaded successfully")
}

func Get() *Config {
	return conf
}
