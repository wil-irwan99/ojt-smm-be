package main

import (
	"project-ojt/delivery"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/config.env")
	if err != nil {
		panic(err)
	}
	delivery.Server().Run()
}
