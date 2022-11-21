package config

import (
	"fmt"
	"os"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type PRTGConfig struct {
	Ip       string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	PRTGConfig
	DbConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	ip := os.Getenv("IP_PRTG")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	c.ApiConfig = ApiConfig{Url: api}
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.PRTGConfig = PRTGConfig{
		Ip:       ip,
		User:     user,
		Password: password,
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
