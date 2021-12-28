package main

import (
	"log"
	"main/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	addr := ":8080"
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	log.Fatalf("Error occurred: %s", err.Error())
	// }
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := database.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	// httpHandler := handler.NewHandler(database)
	// server := &http.Server{
	// 	Handler: httpHandler,
	// }
	// go func() {
	// 	server.Serve(listener)
	// }()
	// defer Stop(server)
	log.Printf("Started server on %s", addr)
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// log.Println(fmt.Sprint(<-ch))
	// log.Println("Stopping API server.")
}
