package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TestEnvKey string
}

var ENV = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		TestEnvKey: getEnv("TEST_ENV_KEY", "some_env_key"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
