package random

import (
	"github.com/acnagy/chaos-scheduler/threads"
)

func Run(thpool []threads.Thread, thpoolSize int, done chan bool) {
	set_priorities(thpool, thpoolSize)
	threads.Work("random", thpool, thpoolSize)

	done <- true
}
