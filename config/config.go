package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv fetching data from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
}

// GetToken return the bot token from .env file
func GetToken() string {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN empty or not found in .env")
	}
	return token
}

// GetOwnerID returning id of the owner from .env file
func GetOwnerID() int64 {
	id := os.Getenv("OWNER_ID")
	if id == "" {
		log.Fatal("OWNER_ID is not found in .env")
	}
	return parseInt(id) // Convert into int64
}

// Just a function to make sure that the string is a number
func parseInt(value string) int64 {
	var result int64
	_, err := fmt.Sscanf(value, "%d", &result)
	if err != nil {
		log.Fatalf("error while parsing: %v", err)
	}
	return result
}
