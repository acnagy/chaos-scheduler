package threads

import (
	"log"
	"math/rand"
	"time"
)

type Thread struct {
	Id       uint16
	Priority uint16
	Worktime int
}

var currentThreadId = uint16(1000) // init w four digits for convenience w weather scheduling

func CreateThread() Thread {
	th := Thread{newThreadId(), 0, rand.Intn(2E9) + 1E6}
	return th
}

func CreateThreadRandomly(ch1 chan Thread, ch2 chan Thread,
	ch3 chan Thread, ch4 chan Thread, thRemaining int, done chan bool) {
	rand.Seed(time.Now().UnixNano())

	for thRemaining > 0 {
		t := rand.Uint32() / 1E9
		time.Sleep(time.Duration(t) * time.Nanosecond)
		log.Printf("[all] new thread! Remaining: %d", thRemaining)
		th := CreateThread()
		ch1 <- th
		ch2 <- th
		ch3 <- th
		ch4 <- th
		thRemaining = thRemaining - 1
	}
	close(ch1)
	close(ch2)
	close(ch3)
	close(ch4)

	done <- true
}

func PickUpThreads(thpool []Thread, maxThreads int, waitingTh chan Thread) []Thread {
	for i := 0; i < cap(thpool); i++ {
		if thpool[i].Id == 0 {
			th := <-waitingTh
			thpool[i] = th
		}
	}

	return thpool
}

func sortThreads(thpool []Thread) (thPoolSorted []Thread) {
	// Bubble sort polices according to priority
	for i := 0; i < len(thpool); i++ {
		for j := 0; j < len(thpool)-1; j++ {
			if thpool[j].Priority > thpool[j+1].Priority {

				temp := thpool[j]
				thpool[j] = thpool[j+1]
				thpool[j+1] = temp
			}
		}
	}
	return thpool
}

func Work(policy string, thpool []Thread) ([]Thread, time.Duration) {
	start := time.Now()
	threadpool := sortThreads(thpool)
	log.Printf("[%s] thread batch sorted by priority\n", policy)

	for i := 0; i < len(thpool); i++ {
		if threadpool[i].Id == 0 {
			log.Printf("[%s] id: %d - working %d ms...",
				policy, threadpool[i].Id, threadpool[i].Worktime/1E6)

			time.Sleep(time.Duration(threadpool[i].Worktime) * time.Nanosecond)

			log.Printf("[%s] id: %d - done\n", policy, threadpool[i].Id)

			threadpool[i] = Thread{0, 0, 0}
		}
	}

	log.Printf("[%s] duration: ~%s\n", policy, time.Since(start).String())
	return threadpool, time.Since(start)
}

func newThreadId() uint16 {
	currentThreadId++
	return currentThreadId
}

func LogThreadpool(policy string, thpool []Thread) {
	log.Printf("[%s] - threadpool:  %+v\n", policy, thpool)
}

func InitWaitingThreads(ch1 chan Thread, ch2 chan Thread,
	ch3 chan Thread, ch4 chan Thread, thpoolSize int) {
	for i := 0; i < thpoolSize; i++ {
		th := CreateThread()
		ch1 <- th
		ch2 <- th
		ch3 <- th
		ch4 <- th
	}
}

func InitThreadpoolControl() []Thread {
	return []Thread{
		{Id: 6, Priority: 2, Worktime: 1000000},
		{Id: 2, Priority: 4, Worktime: 2500000},
		{Id: 4, Priority: 1, Worktime: 50000},
		{Id: 8, Priority: 7, Worktime: 350000},
		{Id: 5, Priority: 8, Worktime: 500},
	}
}
