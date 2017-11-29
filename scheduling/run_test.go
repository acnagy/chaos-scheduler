package scheduling

import (
	"github.com/acnagy/chaos-scheduler/simulation"
	"github.com/acnagy/chaos-scheduler/threads"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const threadsToSimulate = 5
const threadpoolSize = 5

func Benchmark_Random(b *testing.B) {
	disableLogs()

	for n := 0; n < b.N; n++ {
		maxThreads := threadsToSimulate
		thpoolSize := threadpoolSize
		threadsRemaining := maxThreads - thpoolSize
		th := make(chan threads.Thread, thpoolSize)

		threads.InitWaitingThreadsControl(th)
		simulation.CreateThreadRandomly(
			th,
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			threadsRemaining,
			make(chan bool, 1))

		Run("random", thpoolSize, maxThreads, th, make(chan bool, 1))
	}
}

func Benchmark_WeatherStatic(b *testing.B) {
	disableLogs()

	for n := 0; n < b.N; n++ {
		maxThreads := threadsToSimulate
		thpoolSize := threadpoolSize
		threadsRemaining := maxThreads - thpoolSize
		th := make(chan threads.Thread, thpoolSize)

		threads.InitWaitingThreadsControl(th)
		simulation.CreateThreadRandomly(
			th,
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			threadsRemaining,
			make(chan bool, 1))

		Run("weather - static", thpoolSize, maxThreads, th, make(chan bool, 1))
	}
}

func Benchmark_WeatherVariable(b *testing.B) {
	disableLogs()

	for n := 0; n < b.N; n++ {
		maxThreads := threadsToSimulate
		thpoolSize := threadpoolSize
		threadsRemaining := maxThreads - thpoolSize
		th := make(chan threads.Thread, thpoolSize)

		threads.InitWaitingThreadsControl(th)
		simulation.CreateThreadRandomly(
			th,
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			threadsRemaining,
			make(chan bool, 1))

		Run("weather - variable", thpoolSize, maxThreads, th, make(chan bool, 1))
	}
}

func Benchmark_ShortestJobFirst(b *testing.B) {
	disableLogs()

	for n := 0; n < b.N; n++ {
		maxThreads := threadsToSimulate
		thpoolSize := threadpoolSize
		threadsRemaining := maxThreads - thpoolSize
		th := make(chan threads.Thread, thpoolSize)

		threads.InitWaitingThreadsControl(th)
		simulation.CreateThreadRandomly(
			th,
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			make(chan threads.Thread, thpoolSize),
			threadsRemaining,
			make(chan bool, 1))

		Run("sjf", thpoolSize, maxThreads, th, make(chan bool, 1))
	}
}

func disableLogs() {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)
}
