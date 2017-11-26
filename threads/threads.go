package threads

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Thread struct {
	Id       uint16
	Priority uint16
	Worktime int
}

func InitThreadpool(thpool []Thread, thpoolSize int) []Thread {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < thpoolSize; i++ {
		th := Thread{threadId(), 0, rand.Intn(2E9) + 1E6}
		thpool[i] = th

	}

	return thpool
}

func LogThreadpool(thpool []Thread, thpoolSize int) {
	var buffer bytes.Buffer

	for i := 0; i < thpoolSize; i++ {
		th := fmt.Sprintf("id: %5d - priority: %5d - worktime: %5d\n",
			thpool[i].Id,
			thpool[i].Priority,
			thpool[i].Worktime,
		)
		buffer.WriteString(th)
	}
}

func Work(process string, thpool []Thread, thpoolSize int) (runtime int, duration time.Duration) {

	start := time.Now()
	thpool = sortThreads(thpool, thpoolSize)
	LogThreadpool(thpool, thpoolSize)
	for i := 0; i < thpoolSize; i++ {
		log.Printf("[%s] id: %d - working %d ms...",
			process, thpool[i].Id, thpool[i].Worktime/1E6)
		time.Sleep(time.Duration(thpool[i].Worktime) * time.Nanosecond)
		runtime = runtime + thpool[i].Worktime
		log.Printf("[%s] id: %d - done\n", process, thpool[i].Id)

		thpool[i] = Thread{0, 0, 0}
	}

	log.Printf("[%s] total worktime: %d ms, duration: ~%s\n",
		process, runtime/1E6, time.Since(start).String())
	return runtime, time.Since(start)

}

func sortThreads(thpool []Thread, thpoolSize int) (thPoolSorted []Thread) {

	// Bubble sort processes according to priority
	for i := 0; i < thpoolSize; i++ {
		for j := 0; j < thpoolSize-1; j++ {
			if thpool[j].Priority > thpool[j+1].Priority {

				temp := thpool[j]
				thpool[j] = thpool[j+1]
				thpool[j+1] = temp
			}
		}
	}

	return thpool

}

func threadId() (id uint16) {

	rand.Seed(time.Now().UnixNano())
	return uint16(rand.Uint32())
}

func threadWorktime() (worktime uint32) {

	rand.Seed(time.Now().UnixNano())
	return rand.Uint32()

}

func InitThreadpoolControl() []Thread {
	return []Thread{
		{Id: 6, Priority: 2, Worktime: 1000000},
		{Id: 1, Priority: 10, Worktime: 1500000},
		{Id: 7, Priority: 3, Worktime: 3000000},
		{Id: 2, Priority: 4, Worktime: 2500000},
		{Id: 9, Priority: 5, Worktime: 15000},
		{Id: 3, Priority: 9, Worktime: 2000000},
		{Id: 10, Priority: 6, Worktime: 75000},
		{Id: 4, Priority: 1, Worktime: 50000},
		{Id: 8, Priority: 7, Worktime: 10000},
		{Id: 5, Priority: 8, Worktime: 500},
	}
}
