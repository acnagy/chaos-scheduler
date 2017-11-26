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
	threadpool := make([]threads.Thread, *threadpoolSize)
	threadpool = threads.InitThreadpool(threadpool)
	threads.LogThreadpool(threadpool)
	runConfig := fmt.Sprintf("[main] modes: random: %t, weather: %t, sjf: %t; threadpool size: %d\n",
		*randomMode, *weatherMode, *sjfMode, *threadpoolSize)

	log.Printf(runConfig)
	fmt.Printf(runConfig)

	// Run modes concurrently
	randomDone := make(chan bool, 1)
	weatherDone := make(chan bool, 1)
	sjfDone := make(chan bool, 1)

	if *randomMode {
		go scheduling.Run("random", threadpool, randomDone)
	}
	if *weatherMode {
		go scheduling.Run("weather", threadpool, weatherDone)
	}

	if *sjfMode {
		go scheduling.Run("sjf", threadpool, sjfDone)
	}

	<-randomDone
	<-weatherDone
	<-sjfDone

	log.Println("[main] - complete!")

}
