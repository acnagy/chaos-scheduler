package scheduling

import (
	"github.com/acnagy/chaos-scheduler/threads"
)

func sjf_priorities(thpool []threads.Thread) []threads.Thread {

	for i := 0; i < len(thpool); i++ {
		thpool[i].Priority = uint16(thpool[i].Worktime)
	}

	return thpool

}
