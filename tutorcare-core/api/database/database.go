package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST = "db"
	PORT = 5433
)

type Database struct {
	Conn *gorm.DB
}

var ErrNoMatch = fmt.Errorf("Error: no matching table record")

var ErrDuplicate = fmt.Errorf("Error: table record already exists")

var ErrSameUser = fmt.Errorf("Error: two users are the same")

func InitializeDatabase(username, password, database string) (Database, error) {
	db := Database{}
	dsn := "postgres://user:password@db/tutorcare_core?sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	db.Conn = conn
	log.Println("Database connection established")
	return db, nil
}
