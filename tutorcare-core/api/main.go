package main

import (
	"log"
	"main/database"
	"main/handler"
	"os"

	"gopkg.in/olahol/melody.v1"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	m := melody.New()
	dbUserName, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := database.InitializeDatabase(dbUserName, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	httpHandler := handler.RouteHandler(database, m)
	listener := httpHandler.Run(":8080")
	log.Fatal(listener)
	log.Printf("Started server on %s", "8080")
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// log.Println(fmt.Sprint(<-ch))
	// log.Println("Stopping API server.")
}
