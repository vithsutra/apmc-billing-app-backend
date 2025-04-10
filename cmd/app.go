package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vsynclabs/billsoft/pkg/database"
)

func init() {
	// Get the SERVER_MODE from the environment variable
	serverMode := os.Getenv("SERVER_MODE")

	// Check if SERVER_MODE is "dev"  //  SERVER_MODE=dev go run ./cmd
	if serverMode == "dev" {
		// Load the .env file
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Unable to load .env file: %v", err)
		}
		log.Println(".env file loaded successfully.")

	} else if serverMode != "prod" {
		// If SERVER_MODE is neither "dev" nor "prod", show an error and exit
		log.Fatalln("Invalid SERVER_MODE. Please set SERVER_MODE to 'dev' or 'prod'.")
	}
}

func Start() {
	// Load DB connection
	conn := NewDatabaseConnection()
	conn.CheckStatus()
	defer conn.Close()

	// Init DB
	query := database.NewQuery(conn.db)
	if err := query.InitilizeDatabase(); err != nil {
		log.Fatalf("Unable to initilize database: %v", err)
	}

	// Load RabbitMQ
	rabbitmqConn := NewRabbitmqConnection()
	defer rabbitmqConn.conn.Close()
	defer rabbitmqConn.chann.Close()

	// New Server
	server := &http.Server{
		Addr:    os.Getenv("ADDRESS"),
		Handler: NewRouter(conn, rabbitmqConn).mux, // pass it here!
	}
	log.Printf("Server is running on address %v", os.Getenv("ADDRESS"))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start the server: %v", err)
	}
}
