package server

import (
	"log"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load("local-config.env")
	if err != nil {
		log.Fatal("Error loading in enviornment file!")
		return
	}

	// TODO
}