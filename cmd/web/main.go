package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kokweikhong/go-billplz/internal/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := web.New()
	app.Initialize()
	if err := app.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
