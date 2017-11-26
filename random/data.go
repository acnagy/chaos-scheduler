package random

import (
	"github.com/acnagy/chaos-scheduler/threads"
	"math/rand"
	"time"
)

func set_priorities(thpool []threads.Thread, thpool_size int) []threads.Thread {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < thpool_size; i++ {
		thpool[i].Priority = uint16(rand.Int31n(2 * int32(thpool_size)))
	}

	return thpool
}
