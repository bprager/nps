
package nps

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createSchema() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get the connection info
	dbUser := os.Getenv("PG_USER")
	dbPwd := os.Getenv("PG_PW")
	// connect
	dbinfo := fmt.Sprintf("user=%s password=%s sslmode=disable", dbUser, dbPwd)
	db, err := sqlx.Connect("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	// check if table exist
	var sqlStatement = "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'surveys')"
	row := db.QueryRow(sqlStatement)
	var exist bool
	err = row.Scan(&exist)

	if exist {
		log.Print("schema exist")
	} else {
		path := filepath.Join("../backend/", "schema.sql")
		_, err = sqlx.LoadFile(db, path)
		checkErr(err)
		log.Print("schema initialized")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
