package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	DB_Port  int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWT_SECRET  string
	DBConfig    DBConfig
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

	//DB
	host := os.Getenv("DB_HOST")
	if host == "" {
		fmt.Println("DB HOST is not set")
		os.Exit(1)
	}

	port := os.Getenv("DB_PORT")
	int_port, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("DB_PORT is not set")
		os.Exit(1)
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		fmt.Println("DB_USER is not set")
		os.Exit(1)
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Println("DB PASSWORD is not set")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB_NAME is not set")
		os.Exit(1)
	}

	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		fmt.Println("DB_SSLMODE is not set")
		os.Exit(1)
	}

	dbConfig := DBConfig{
		Host:     host,
		DB_Port:  int(int_port),
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}

	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(httpPortInt),
		JWT_SECRET:  jwtSecret,
		DBConfig:    dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
