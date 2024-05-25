package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	RedisUrl string
}

type ConfigError struct {
	Field string
}

func (e ConfigError) Error() string {
	return fmt.Sprintf("Missing env var: %s", e.Field)
}

func FromEnv() (Config, error) {
	// return Config{}, nil
	host, present := os.LookupEnv("TODOLIST_HOST")
	if !present {
		return Config{}, ConfigError{"TODOLIST_HOST"}
	}

	port, present := os.LookupEnv("TODOLIST_PORT")
	if !present {
		return Config{}, ConfigError{"TODOLIST_PORT"}
	}

	redisUrl, present := os.LookupEnv("TODOLIST_REDIS_URL")
	if !present {
		return Config{}, ConfigError{"TODOLIST_REDIS_URL"}
	}

	return Config{host, port, redisUrl}, nil
}
