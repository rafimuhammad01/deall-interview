package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rafimuhammad01/api-gateway/api"
	"log"
	"os"
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
