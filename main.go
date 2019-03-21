package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := getDB("root:rootpass@tcp(0.0.0.0:3306)/blog")
	if err != nil {
		log.Fatalf("Failed getting DB. Error :%s", err)
	}
	defer db.Close()

	if err = prepareStmts(db); err != nil {
		log.Fatalf("Failed preparing al statements. Error :%s", err)
	}
	defer closeStmts()

	addPosts()
}

func getDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return db, err
	}
	return db, nil
}
