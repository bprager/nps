package nps

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	log.Print("init")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get the connection info
	dbUser := os.Getenv("PG_USER")
	dbPwd := os.Getenv("PG_PW")
	dbHost := os.Getenv("PG_HOST")
	// connect
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", dbHost, dbUser, dbPwd)
	db, err = sqlx.Connect("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	// check if table exist

	db, err = sqlx.Open("postgres", dbinfo)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	tables := []string{}
	db.Select(&tables, "SHOW TABLES")
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
