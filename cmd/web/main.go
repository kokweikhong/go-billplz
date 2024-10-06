package main

import (
	"fmt"

	"github.com/kokweikhong/go-billplz/internal/web"
)

func main() {
	app := web.New()
	app.InitRoutes()
	if err := app.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
