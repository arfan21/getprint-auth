package main

import (
	"log"
	"os"

	"github.com/arfan21/getprint-service-auth/config"
	"github.com/arfan21/getprint-service-auth/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	kidToken := os.Getenv("JWT_KID")
	kidRefreshToken := os.Getenv("REFRESH_TOKEN_KID")

	if err := config.CreateKey(kidToken, "token"); err != nil {
		log.Println(err)
	}
	if err := config.CreateKey(kidRefreshToken, "refreshToken"); err != nil {
		log.Println(err)
	}

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
