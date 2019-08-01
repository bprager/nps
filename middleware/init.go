package nps

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	// Postgresql driver
	_ "github.com/lib/pq"
)

// DB database handler
var DB *sqlx.DB

func init() {
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
	DB, err = sqlx.Open("postgres", dbinfo)
	checkErr(err)
	defer DB.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
