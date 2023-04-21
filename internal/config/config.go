package config

import "os"

// Config is a struct that holds all the configuration variables
type Config struct {
	Database struct {
		Url string
	}
	Server struct {
		Port string
	}
	Env string
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	return &Config{
		Database: struct {
			Url string
		}{
			Url: os.Getenv("DATABASE_URL"),
		},
		Server: struct {
			Port string
		}{
			Port: getPort(),
		},
		Env: getEnv(),
	}
}

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}

	return env
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
