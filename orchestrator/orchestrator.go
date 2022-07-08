package orchestrator

import (
	"fmt"

	"github.com/dboreham/db-locking-test/sessions"
)

func Run() {
	fmt.Println("Orchestrator run called")
	sessions.Run()
}
