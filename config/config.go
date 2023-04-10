package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	ServerHost string
}

var (
	conf *Config
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env is not loaded properly")
	}

	conf = new(Config)
	conf.ServerPort, _ = strconv.Atoi(os.Getenv(`PORT`))
	conf.ServerHost = os.Getenv(`HOST`)
}

func Get() *Config {
	return conf
}
