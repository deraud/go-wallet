package main

import (
	"log"
	"main/app"
)

func main() {
	application := app.NewApp()
	if err := application.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
