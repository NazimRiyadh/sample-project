package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWT_SECRET  string
}

var configurations *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Eror while loading env variables", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION is not set")
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME is not set")
	}

	httpPort := os.Getenv("HTTP_PORT")
	httpPortInt, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("HTTP_PORT is not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT secret not aailable")
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(httpPortInt),
		JWT_SECRET:  jwtSecret,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
