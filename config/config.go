package config

import "os"

type ApiConfig struct {
	Url string
}

type PRTGConfig struct {
	Ip       string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	PRTGConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	ip := os.Getenv("IP_PRTG")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	c.ApiConfig = ApiConfig{Url: api}

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
