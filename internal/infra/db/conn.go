package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "alpha.mkdb.sh"
	port     = 5432
	user     = "kbslymyv"
	password = "qhcizqlxyxvxjmgqgvcn"
	dbname   = "bggosigm"
)

func ConnectDb() (*sql.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
