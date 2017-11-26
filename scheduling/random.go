package scheduling

import (
	"github.com/acnagy/chaos-scheduler/threads"
	"math/rand"
	"time"
)

func random_priorities(thpool []threads.Thread) []threads.Thread {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(thpool); i++ {
		thpool[i].Priority = uint16(rand.Int31n(2 * int32(len(thpool))))
	}

	return thpool
}
