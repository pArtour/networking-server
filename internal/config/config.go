package config

import "os"

var Cfg *Config

// Config is a struct that holds all the configuration variables
type Config struct {
	Database struct {
		Url string
	}
	Server struct {
		Port string
	}
	Env       string
	JWTSecret string
}

// InitConfig initializes the configuration variables
func InitConfig() {
	Cfg = &Config{
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
		Env:       getEnv(),
		JWTSecret: os.Getenv("JWT_SECRET"),
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
