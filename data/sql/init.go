package sql

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // adds the mysql driver
)

// DB manages communication to the database
type DB struct {
	common service
	db     *gorm.DB
	DSN    string

	// Services
	User *UserService
}

type service struct {
	db *DB
}

// NewConnection establishes the sql connection
func NewConnection(
	DSN string,
	debug bool,
) *DB {
	db, err := gorm.Open("mysql", DSN)
	if err != nil {
		panic("Failed to connect to database, err:" + err.Error())
	}

	log.Printf("Database connection established")

	if debug {
		db.LogMode(debug)
	}

	connection := &DB{
		db:  db,
		DSN: DSN,
	}

	return connection
}
