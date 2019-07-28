package nps

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

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
	DB, err = sqlx.Connect("postgres", dbinfo)
	checkErr(err)
	defer DB.Close()
	// check if table exist

	tables := []string{}
	DB.Select(&tables, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname='public';")
	log.Printf("Number of tables found: %d", len(tables))
	log.Printf("%v", tables)
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
