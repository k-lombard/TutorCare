package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	HOST = "db"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

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
	database, err := Initialize(dbUser, dbPassword, dbName)
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

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := "postgres://user:password@db/tutorcare_core?sslmode=disable"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

// func Stop(server *http.Server) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := server.Shutdown(ctx); err != nil {
// 		log.Printf("Could not shut down server correctly: %v\n", err)
// 		os.Exit(1)
// 	}
// }
