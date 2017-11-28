package scheduling

import (
	"fmt"
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"os"
)

func Run(policy string, thpoolSize int, maxThreads int, waitingThreads chan threads.Thread, done chan bool) {
	stout := log.New(os.Stdout, "[chaos-scheduler]", log.Ldate|log.Ltime|log.Lshortfile)
	startedMsg := fmt.Sprintf("[%s] - started", policy)

	log.Println(startedMsg)
	stout.Println(startedMsg)

	for {
		threadpool := make([]threads.Thread, thpoolSize)
		threads.PickUpThreads(threadpool, maxThreads, waitingThreads)
		threads.LogThreadpool(policy, threadpool)

		switch policy {
		case "random":
			threadpool = random_priorities(threadpool)
		case "weather - static", "weather - variable":
			threadpool = weather_priorities(policy, threadpool)
		case "sjf":
			threadpool = sjf_priorities(threadpool)
		case "control": // for benchmarking
		default:
			stout.Panicln("policy selection error")
			log.Panicln("invalid policy selection string")
		}

		threads.Work(policy, threadpool)
		log.Printf("[%s] work complete", policy)

		_, ok := <-waitingThreads
		log.Printf("[%s] threads channel status: %t\n", policy, ok)
		threads.LogThreadpool(policy, threadpool)

		if !ok {
			done <- true
			break
		}
	}

	doneMsg := fmt.Sprintf("[%s] - DONE", policy)
	log.Println(doneMsg)
	stout.Println(doneMsg)
}
