package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AhGr3y/my-grocery-list/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	DB *database.Queries
}

func main() {

	// Load environment variables in .env
	godotenv.Load()

	// Get database url from .env
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Initialize appConfig
	dbQueries := database.New(db)
	cfg := config{
		DB: dbQueries,
	}

	runApp(&cfg)
}
