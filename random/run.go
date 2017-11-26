package random

import (
	"github.com/acnagy/chaos-scheduler/threads"
)

func Run(thpool []threads.Thread, thpool_size int, done chan bool) {
	set_priorities(thpool, thpool_size)
	//threads.Print_Threadpool(thpool, thpool_size)
	threads.Work("random", thpool, thpool_size)

	done <- true
}
