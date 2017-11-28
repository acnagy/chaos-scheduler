package scheduling

import (
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
)

func sjf_priorities(thpool []threads.Thread) []threads.Thread {

	for i := 0; i < len(thpool); i++ {
		thpool[i].Priority = uint16(thpool[i].Worktime)
	}

	log.Println("[sjf] thread batch prioritized")

	return thpool

}
