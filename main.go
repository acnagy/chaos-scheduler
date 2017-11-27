package main

import (
	"flag"
	"fmt"
	"github.com/acnagy/chaos-scheduler/scheduling"
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"os"
)

var threadpoolSize = flag.Int("t", 5, "set size for threadpool")
var numberThreads = flag.Int("n", 10, "set number of threads created during program execution")
var randomMode = flag.Bool("r", true, "run random scheduling")
var weatherMode = flag.Bool("w", true, "run weather scheduling")
var sjfMode = flag.Bool("s", true, "run shortest-job-first scheduling")

func main() {
	fmt.Print("chaos-scheduler starting... ")

	// Setup logging
	logfile, err := os.OpenFile("dev/log.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatalf("Couldn't open log file: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Printf("[main] Started!")

	//artifact := log.New(os.OpenFile("dev/artifact.txt", "[main]", log.Ldate|log.Ltime|log.Lshortfile))

	// Determine/Notify running modes
	flag.Parse()
	runConfig := fmt.Sprintf(
		"[main] modes: random: %t, weather: %t, sjf: %t; threadpool size: %d, max threads: %d\n",
		*randomMode, *weatherMode, *sjfMode, *threadpoolSize, *numberThreads,
	)
	log.Printf(runConfig)
	fmt.Printf(runConfig)

	// Run modes concurrently
	randomDone := make(chan bool, 1)
	randomThr := make(chan threads.Thread, *numberThreads)

	weatherDone := make(chan bool, 1)
	weatherThr := make(chan threads.Thread, *numberThreads)

	sjfDone := make(chan bool, 1)
	sjfThr := make(chan threads.Thread, *numberThreads)

	threadsDone := make(chan bool, 1)
	threads.InitWaitingThreads(randomThr, weatherThr, sjfThr, *threadpoolSize)
	if *threadpoolSize != *numberThreads {
		go threads.CreateThreadRandomly(
			randomThr, weatherThr, sjfThr,
			*numberThreads-*threadpoolSize,
			threadsDone,
		)
	}

	if *randomMode {
		go scheduling.Run("random", *threadpoolSize, randomThr, randomDone)
	}
	if *weatherMode {
		go scheduling.Run("weather", *threadpoolSize, weatherThr, weatherDone)
	}
	if *sjfMode {
		go scheduling.Run("sjf", *threadpoolSize, sjfThr, sjfDone)
	}

	<-randomDone
	<-weatherDone
	<-sjfDone
	<-threadsDone

	log.Println("[main] - complete!")
}
