package scheduling

import (
	"github.com/acnagy/chaos-scheduler/threads"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkRun_Control(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	for n := 0; n < b.N; n++ {
		thpool := threads.InitThreadpoolControl()
		Run("benchmark - control", thpool, len(thpool), make(chan bool, 1))
	}
}

func BenchmarkRun_Random(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	for n := 0; n < b.N; n++ {
		thpool := threads.InitThreadpoolControl()
		Run("benchmark - control", thpool, len(thpool), make(chan bool, 1))
	}

}
