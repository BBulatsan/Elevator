package db

import (
	"database/sql"

	"Elevator/models"
	_ "github.com/mattn/go-sqlite3"
)

//CREATE TABLE IF NOT EXISTS log (id INTEGER PRIMARY KEY, 'Name' NUMBER , Place NUMBER,QuantityOfPeople Number, ACT message_text);

func Db() *sql.DB {
	db, err := sql.Open("sqlite3", "./modules/db/dbs/log.db")
	if err != nil {
		panic(err)
	}
	return db
}

func AddLog(db *sql.DB, log models.Log) {
	statement, err := db.Prepare("INSERT INTO log ('Name', Place,QuantityOfPeople,ACT ) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(log.Name, log.Place, log.QuantityOfPeople, log.Action)
}
