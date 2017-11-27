package scheduling

import (
	"fmt"
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"os"
)

func Run(policy string, thpoolSize int, waitingThreads chan threads.Thread, done chan bool) {
	stout := log.New(os.Stdout, "[chaos-scheduler]", log.Ldate|log.Ltime|log.Lshortfile)
	startedMsg := fmt.Sprintf("[%s] - started", policy)

	log.Println(startedMsg)
	stout.Println(startedMsg)

	for {
		threadpool := make([]threads.Thread, thpoolSize, thpoolSize)
		threads.PickUpThreads(threadpool, waitingThreads)
		threads.LogThreadpool(policy, threadpool)

		switch policy {
		case "random":
			threadpool = random_priorities(threadpool)
		case "weather":
			threadpool = weather_priorities(threadpool)
		case "sjf":
			threadpool = sjf_priorities(threadpool)
		default:
			stout.Println("policy selection error")
			log.Panicln("invalid policy selection string")
		}

		threads.Work(policy, threadpool)

		_, ok := <-waitingThreads
		log.Printf("[%s] - threads channel status: %t\n", policy, ok)
		threads.LogThreadpool(policy, threadpool)

		if !ok {
			done <- true
			break
		}
	}

	doneMsg := fmt.Sprintf("[%s] - done", policy)
	log.Println(doneMsg)
	stout.Println(doneMsg)
}
