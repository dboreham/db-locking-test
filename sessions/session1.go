package sessions

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func run() {
	connectionString := fmt.Sprintf("postgresql://db:db@localhost:57381/db?sslmode=disable")
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

// Call this to start the session
// Session waits before advancing to execute the next phase
// upon recipt of a message on its advance channel
// Session sends a message on its phase channel to indicate it has completed the phase
// (and is waiting)
func Start(advance <-chan int, phase chan<- int) {
	go func() {
		// Wait on a message on the advance channel
		x := <-advance
		log.Println(x)
		// Execute a sleep query
		log.Println("Executing query")
		run()
		// Signal to the orchestrator that we've completed the query
		phase <- 2
	}()
}
