package config

import (
	"database/sql"
	"fmt"
	"log"	
	"os"
	"strconv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

var (
	
    
	host     = os.Getenv("DB_HOST")
	port     = 5432
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbname   = os.Getenv("DB_NAME")
)

func ConnectDB() {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
    if err != nil {
        dbPort = 5432
    }
    host     = os.Getenv("DB_HOST")
	port     = dbPort
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbname   = os.Getenv("DB_NAME")
	
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	DB = db
	//fmt.Println("Database connected")
	log.Println("Database connected")
}

func CloseDB() {
	DB.Close()
}
