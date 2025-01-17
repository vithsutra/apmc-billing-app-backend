package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vsynclabs/billsoft/pkg/database"
)

func Start() {
	// Loading env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Unable to load env: %v", err)
	}
	// Database Connection
	conn := NewDatabaseConnection()
	conn.CheckStatus()
	defer conn.Close()
	// Initilize Database
	query := database.NewQuery(conn.db)
	if err := query.InitilizeDatabase(); err != nil {
		log.Fatalf("Unable to initilize database: %v", err)
	}
	// New Server
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: NewRouter(conn).mux,
	}
	log.Printf("Server is running on port %v", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start the server: %v", err)
	}
}
