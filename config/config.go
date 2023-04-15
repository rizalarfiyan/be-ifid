package config

import (
	"be-ifid/utils"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  int
	Host  string
	Cors  CorsConfigs
	MQTT  MQTTConfigs
	DB    DBConfigs
	Redis RedisConfigs
	Email EmailConfigs
	FE    FEConfigs
	JWT   JWTConfigs
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
	Name               string
	Host               string
	Port               int
	User               string
	Password           string
	ConnectionIdle     time.Duration
	ConnectionLifetime time.Duration
	MaxIdle            int
	MaxOpen            int
}

type RedisConfigs struct {
	Host            string
	Port            int
	User            string
	Password        string
	ExpiredDuration time.Duration
	DialTimeout     time.Duration
}

type EmailConfigs struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type FEConfigs struct {
	BaseUrl         string
	AuthRedirectUrl string
}

type JWTConfigs struct {
	SecretKey string
	Expired   time.Duration
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
	conf.DB.ConnectionIdle = utils.GetEnvAsTimeDuration("DB_CONNECTION_IDLE", 1*time.Minute)
	conf.DB.ConnectionLifetime = utils.GetEnvAsTimeDuration("DB_CONNECTION_LIFETIME", 5*time.Minute)
	conf.DB.MaxIdle = utils.GetEnvAsInt("DB_MAX_IDLE", 20)
	conf.DB.MaxOpen = utils.GetEnvAsInt("DB_MAX_OPEN", 50)

	conf.Redis.Host = utils.GetEnv("REDIS_HOST", "")
	conf.Redis.Port = utils.GetEnvAsInt("REDIS_PORT", 6379)
	conf.Redis.User = utils.GetEnv("REDIS_USER", "")
	conf.Redis.Password = utils.GetEnv("REDIS_PASSWORD", "")
	conf.Redis.ExpiredDuration = utils.GetEnvAsTimeDuration("REDIS_EXPIRED_DURATION", 15*time.Minute)
	conf.Redis.DialTimeout = utils.GetEnvAsTimeDuration("REDIS_DIAL_TIMEOUT", 5*time.Minute)

	conf.Email.Host = utils.GetEnv("EMAIL_HOST", "")
	conf.Email.Port = utils.GetEnvAsInt("EMAIL_PORT", 587)
	conf.Email.User = utils.GetEnv("EMAIL_USER", "")
	conf.Email.Password = utils.GetEnv("EMAIL_PASSWORD", "")
	conf.Email.From = utils.GetEnv("EMAIL_FROM", "")

	conf.FE.BaseUrl = utils.GetEnv("FE_BASE_URL", "")
	conf.FE.AuthRedirectUrl = utils.GetEnv("FE_AUTH_REDIRECT_URL", "")

	conf.JWT.SecretKey = utils.GetEnv("JWT_SECRET_KEY", "")
	conf.JWT.Expired = utils.GetEnvAsTimeDuration("JWT_EXPIRED", 5*24*time.Hour)

	utils.Success("Config is loaded successfully")
}

func Get() *Config {
	return conf
}
