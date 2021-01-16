package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	pg "gopkg.in/pg.v4"
)

var db *sql.DB

var PG *pg.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgespswd"
	dbname   = "oslash"
)

func DBConnecting() {
	InitPG()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetClient() *sql.DB {
	return db
}

func InitPG() {
	PG = pg.Connect(&pg.Options{
		Addr:     host,
		User:     user,
		Password: password,
		Database: dbname,
	})
}

func GetPGClient() *pg.DB {
	return PG
}
