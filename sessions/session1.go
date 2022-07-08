package sessions

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Run() {
	connectionString := fmt.Sprintf("postgresql://db:db@localhost:49153/db?sslmode=disable")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := db.Query("select 1")
	if err != nil {
		log.Fatalln(err)
	}
	var row string
	for rows.Next() {
		rows.Scan(&row)
		fmt.Println(row)
	}
}
