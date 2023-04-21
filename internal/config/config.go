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
			Port: os.Getenv("PORT"),
		},
	}
}
