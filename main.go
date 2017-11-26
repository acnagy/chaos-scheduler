package main

import (
	"flag"
	"fmt"
	"github.com/acnagy/chaos-scheduler/scheduling"
	"github.com/acnagy/chaos-scheduler/threads"
	"log"
	"os"
)

var threadpoolSize = flag.Int("t", 10, "set size for threadpool")
var weatherMode = flag.Bool("w", true, "run weather scheduling")
var randomMode = flag.Bool("r", true, "run random scheduling")

func main() {
	fmt.Print("chaos-scheduler starting... ")

	// Setup logging
	file, err := os.OpenFile("dev/log.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatalf("Couldn't open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Started!")

	// Determine/Notify running modes
	flag.Parse()
	threadpool := make([]threads.Thread, *threadpoolSize)
	threadpool = threads.InitThreadpool(threadpool, *threadpoolSize)
	threads.LogThreadpool(threadpool, *threadpoolSize)
	runConfig := fmt.Sprintf("[main] modes: weather: %t, random: %t, threadpool size: %d\n",
		*weatherMode, *randomMode, *threadpoolSize)

	log.Printf(runConfig)
	fmt.Printf(runConfig)

	// Run modes concurrently
	randomDone := make(chan bool, 1)
	if *randomMode {
		go scheduling.Run("random", threadpool, *threadpoolSize, randomDone)
	}

	<-randomDone
	log.Println("[main] - complete!")

}
