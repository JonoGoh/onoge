package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Get the token from the environment variables
    token := os.Getenv("APP_ID")
    if token == "" {
        log.Fatalf("TOKEN not set in environment")
    }

    fmt.Println("Token:", token)
}