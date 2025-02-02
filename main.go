package main

import (
	"url-shortner/database"
	"url-shortner/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database.ConnectDB()
}

func main() {
	router.ClientRoutes()
}
