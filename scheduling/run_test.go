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

	th := make(chan threads.Thread, 5)
	threads.InitWaitingThreadsControl(th)

	for n := 0; n < b.N; n++ {
		Run("control", 5, 5, th, make(chan bool, 1))
	}
}

func BenchmarkRun_Random(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	th := make(chan threads.Thread, 5)
	threads.InitWaitingThreadsControl(th)

	for n := 0; n < b.N; n++ {
		Run("random", 5, 5, th, make(chan bool, 1))
	}
}

func BenchmarkRun_WeatherStatic(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	th := make(chan threads.Thread, 5)
	threads.InitWaitingThreadsControl(th)

	for n := 0; n < b.N; n++ {
		Run("weather - static", 5, 5, th, make(chan bool, 1))
	}
}

func BenchmarkRun_WeatherVariable(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	th := make(chan threads.Thread, 5)
	threads.InitWaitingThreadsControl(th)

	for n := 0; n < b.N; n++ {
		Run("weather - static", 5, 5, th, make(chan bool, 1))
	}
}

func BenchmarkRun_ShortestJobFirst(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)

	th := make(chan threads.Thread, 5)
	threads.InitWaitingThreadsControl(th)

	for n := 0; n < b.N; n++ {
		Run("sjf", 5, 5, th, make(chan bool, 1))
	}
}
