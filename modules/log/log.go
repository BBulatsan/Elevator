package log

import (
	"database/sql"
	"fmt"

	"Elevator/models"
	"Elevator/modules/db"
)

func Log(con *sql.DB, ch chan models.Log) {
	for i := 1; ; i++ {
		log, isOk := <-ch
		if !isOk {
			break
		}
		db.AddLog(con, log)
		fmt.Println("id:", i, "Elevator:", log.Name, "Place:", log.Place, "Num of People:", log.QuantityOfPeople, "Action:", log.Action)
	}
}
