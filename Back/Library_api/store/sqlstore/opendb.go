package sqlstore

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "172.19.0.3"
	port     = 5432
	user     = "postgres"
	password = "1234qwer!"
	dbname   = "postgres"
)

func DBopen() (db *sql.DB) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected!")

	return db
}
