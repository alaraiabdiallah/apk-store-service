package main

import (
	"github.com/alaraiabdiallah/apk-store-service/data_sources/mongods"
	httpProtocol "github.com/alaraiabdiallah/apk-store-service/protocols/http"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	api_key := os.Getenv("API_KEY")
	if api_key == "" {
		log.Fatal("API_KEY ENVIROMENT variable not defined")
		os.Exit(3)
	}
	mongods.CheckConnection()
	httpProtocol.Run()
}
