package config

import "os"

type Config struct {
	DSN 		string
	HTTPPort 	string
	GinMode 	string
}

func Load() Config {
	return Config{
		DSN: getEnv("TODO_APP_DSN", "todo.db"),
		HTTPPort: getEnv("TODO_APP_HTTP_PORT", ":8080"),
		GinMode: getEnv("TODO_APP_GIN_MODE", "debug"),
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}