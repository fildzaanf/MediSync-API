package main

import (
	"app/config"
	"app/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Routes()

	// start the server on port 8000
	e.Logger.Fatal(e.Start(":8000"))
}