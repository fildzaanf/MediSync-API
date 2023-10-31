package main

import (
	"app/config"
	"app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}

	config.Init()

	e := routes.Routes()

	// start the server on port 8000
	e.Logger.Fatal(e.Start(":8000"))
}
