package config

import "os"

type Config struct {
	Database struct {
		Url string
	}
	Server struct {
		Port string
	}
}

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
			Port: os.Getenv("PORT"),
		},
	}
}
