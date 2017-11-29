package simulation

import (
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"math/rand"
	"time"
)

func CreateThreadRandomly(ch1 chan threads.Thread, ch2 chan threads.Thread,
	ch3 chan threads.Thread, ch4 chan threads.Thread, threadsRemaining int, noMoreThreads chan bool) {
	rand.Seed(time.Now().UnixNano())

	// if the simulation is setup with more threads, randomly generate them
	// and queue up in the scheduling tasks' channels
	for threadsRemaining > 0 {
		t := rand.Uint32() / 1E9
		time.Sleep(time.Duration(t) * time.Nanosecond)
		log.Printf("[all] new thread! Remaining: %d", threadsRemaining)
		th := threads.CreateThread()
		ch1 <- th
		ch2 <- th
		ch3 <- th
		ch4 <- th
		threadsRemaining = threadsRemaining - 1
	}

	// indicate there are no more new threads to simulate
	noMoreThreads <- true
	close(noMoreThreads)
}
