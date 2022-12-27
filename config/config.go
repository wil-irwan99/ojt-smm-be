package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
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

func Loadenv(key string) string {

	viper.SetConfigFile("config/config.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file%s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion %s", key)
	}

	return value
}

func (c *Config) readConfig() {
	api := Loadenv("API_URL")
	ipBIB := Loadenv("IP_BIB")
	ipKIM := Loadenv("IP_KIM")
	ipMAL := Loadenv("IP_MAL")
	ipBSL := Loadenv("IP_BSL")
	ipSML := Loadenv("IP_SML")
	ipMSIG := Loadenv("IP_MSIG")
	ipBCHO := Loadenv("IP_BCHO")
	user := Loadenv("USER")
	password := Loadenv("PASSWORD")
	userBCHO := Loadenv("USER_BCHO")
	passwordBCHO := Loadenv("PASSWORD_BCHO")
	passwordSML := Loadenv("PASSWORD_SML")

	dbHost := Loadenv("DB_HOST")
	dbPort := Loadenv("DB_PORT")
	dbUser := Loadenv("DB_USER")
	dbPassword := Loadenv("DB_PASSWORD")
	dbName := Loadenv("DB_NAME")

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
