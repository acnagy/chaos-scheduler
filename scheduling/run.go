package scheduling

import (
	"fmt"
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"os"
)

func Run(policy string,
	thpoolSize int,
	maxThreads int,
	waitingThreads chan threads.Thread,
	noMoreThreads chan bool) {

	stout := log.New(os.Stdout, "[chaos-scheduler] ", log.Ldate|log.Ltime|log.Lshortfile)
	startedMsg := fmt.Sprintf("[%s] started", policy)

	log.Println(startedMsg)
	stout.Println(startedMsg)

	simulate := true

	for simulate {
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

		select {
		case th := <-waitingThreads:
			// check if there's a thread to read
			// return to channel if something's there
			waitingThreads <- th
		default:
			noMoreThreads <- true
			noMoreThreadsMsg := fmt.Sprintf("[%s] no more simulated threads", policy)
			log.Println(noMoreThreadsMsg)
			stout.Println(noMoreThreadsMsg)
			simulate = false
		}
	}
}
