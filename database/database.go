package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/pascaloseko/go-rest-fibre-postgres/config"
)

// DB instance
var DB *sql.DB

// Connect func
func Connect() error {
	var err error
	pg := config.Config("DB_PORT")

	port, err := strconv.ParseUint(pg, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	CreateProductTable()
	log.Println("Connection Opened Db")
	return nil
}
