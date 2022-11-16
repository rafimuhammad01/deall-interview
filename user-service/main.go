package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafimuhammad01/user-service/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Default().Println(".env not found")
	}
	srv := api.NewServer()
	srv.Init()
	srv.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
