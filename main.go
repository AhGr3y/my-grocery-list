package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables in .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello world!")
}
