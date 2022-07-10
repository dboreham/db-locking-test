package orchestrator

import (
	"fmt"
	"log"

	"github.com/dboreham/db-locking-test/sessions"
)

func Run() {
	fmt.Println("Orchestrator run called")
	advance := make(chan int)
	phase := make(chan int)
	sessions.Start(advance, phase)
	advance <- 1
	what := <-phase
	log.Println(what)
}
