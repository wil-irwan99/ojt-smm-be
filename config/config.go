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
	IpBIB        string
	IpKIM        string
	IpMAL        string
	IpBSL        string
	IpSML        string
	IpMSIG       string
	IpBCHO       string
	User         string
	Password     string
	UserBCHO     string
	PasswordBCHO string
	PasswordSML  string
}

type Config struct {
	ApiConfig
	PRTGConfig
	DbConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	ipBIB := os.Getenv("IP_BIB")
	ipKIM := os.Getenv("IP_KIM")
	ipMAL := os.Getenv("IP_MAL")
	ipBSL := os.Getenv("IP_BSL")
	ipSML := os.Getenv("IP_SML")
	ipMSIG := os.Getenv("IP_MSIG")
	ipBCHO := os.Getenv("IP_BCHO")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	userBCHO := os.Getenv("USER_BCHO")
	passwordBCHO := os.Getenv("PASSWORD_BCHO")
	passwordSML := os.Getenv("PASSWORD_SML")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	c.ApiConfig = ApiConfig{Url: api}
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.PRTGConfig = PRTGConfig{
		IpBIB:        ipBIB,
		IpKIM:        ipKIM,
		IpMAL:        ipMAL,
		IpBSL:        ipBSL,
		IpSML:        ipSML,
		IpMSIG:       ipMSIG,
		IpBCHO:       ipBCHO,
		User:         user,
		Password:     password,
		UserBCHO:     userBCHO,
		PasswordBCHO: passwordBCHO,
		PasswordSML:  passwordSML,
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
