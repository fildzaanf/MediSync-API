package main

import (
	"app/config"
	"app/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {

	_, err := os.Stat(".env")
    if err == nil {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Failed to fetch .env file")
        }
    }

	config.Init()

	e := routes.Routes()

	// start the server on port 8000 
	port := os.Getenv("PORT")
	setPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(setPort))
}
