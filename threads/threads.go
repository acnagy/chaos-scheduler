package threads

import (
	"fmt"
	"math/rand"
	"time"
)

type Thread struct {
	Id       uint16
	Priority uint16
	Worktime int
}

func Init_Threadpool(thpool []Thread, thpool_size int) []Thread {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < thpool_size; i++ {
		th := Thread{thread_id(), 0, rand.Intn(2E9) + 1E6}
		thpool[i] = th

	}

	return thpool
}

func Print_Threadpool(thpool []Thread, thpool_size int) {
	for i := 0; i < thpool_size; i++ {
		fmt.Printf("id: %5d - priority: %5d - worktime: %5d\n",
			thpool[i].Id,
			thpool[i].Priority,
			thpool[i].Worktime,
		)
	}
	fmt.Println()
}

func Work(process string, thpool []Thread, thpool_size int) {

	// Bubble sort processes according to priority
	for i := 0; i < thpool_size; i++ {
		for j := 0; j < thpool_size-1; j++ {
			if thpool[j].Priority > thpool[j+1].Priority {

				temp := thpool[j]
				thpool[j] = thpool[j+1]
				thpool[j+1] = temp
			}
		}
	}

	Print_Threadpool(thpool, thpool_size)

	for i := 0; i < thpool_size; i++ {
		fmt.Printf("[%s] - working %d ms...", process, thpool[i].Worktime/1E6)
		time.Sleep(time.Duration(thpool[i].Worktime) * time.Nanosecond)
		fmt.Printf("done\n")
	}

}

func thread_id() (id uint16) {

	rand.Seed(time.Now().UnixNano())
	return uint16(rand.Uint32())
}

func thread_worktime() (worktime uint32) {

	rand.Seed(time.Now().UnixNano())
	return rand.Uint32()

}
