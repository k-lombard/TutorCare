package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "db"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

var ErrNoMatch = fmt.Errorf("no matching record")

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
