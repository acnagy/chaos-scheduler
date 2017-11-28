package scheduling

import (
	"github.com/acnagy/chaos-scheduler/threads"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkRun_Random(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	for n := 0; n < b.N; n++ {
		//thpool := threads.InitThreadpoolControl()
		Run("random", 10, 10, make(chan threads.Thread, 10), make(chan bool, 1))
	}
}

func BenchmarkRun_WeatherStatic(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	for n := 0; n < b.N; n++ {
		//thpool := threads.InitThreadpoolControl()
		Run("weather - static", 10, 10, make(chan threads.Thread, 10), make(chan bool, 1))
	}
}

func BenchmarkRun_WeatherVariable(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	for n := 0; n < b.N; n++ {
		//thpool := threads.InitThreadpoolControl()
		Run("weather - static", 10, 10, make(chan threads.Thread, 10), make(chan bool, 1))
	}
}

func BenchmarkRun_ShortestJobFirst(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	for n := 0; n < b.N; n++ {
		//thpool := threads.InitThreadpoolControl()
		Run("sjf", 10, 10, make(chan threads.Thread, 10), make(chan bool, 1))
	}
}
