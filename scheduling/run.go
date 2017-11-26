package scheduling

import (
	"fmt"
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"os"
)

func Run(process string, thpool []threads.Thread, thpoolSize int, done chan bool) {
	stout := log.New(os.Stdout, "[chaos-scheduler]", log.Ldate|log.Ltime|log.Lshortfile)

	startedMsg := fmt.Sprintf("[%s] - started", process)
	log.Println(startedMsg)
	stout.Println(startedMsg)

	switch process {
	case "random":
		random_priorities(thpool, thpoolSize)
	default:
		stout.Println("process selection error")
		log.Panicln("invalid process selection string")
	}

	threads.Work(process, thpool, thpoolSize)

	doneMsg := fmt.Sprintf("[%s] - done", process)
	log.Println(doneMsg)
	stout.Println(doneMsg)

	done <- true
}
