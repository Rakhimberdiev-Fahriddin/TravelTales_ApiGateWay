package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT        string
	AUTH_PORT        string
	CONTENT_PORT     string
	ACCESS_TOKEN     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	cfg := Config{}
	cfg.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", ":8080"))
	cfg.AUTH_PORT = cast.ToString(Coalesce("AUTH_PORT", ":50051"))
	cfg.CONTENT_PORT = cast.ToString(Coalesce("CONTENT_PORT", ":50053"))
	cfg.ACCESS_TOKEN = cast.ToString(Coalesce("ACCESS_TOKEN", "my_secret_key"))

	return cfg
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
