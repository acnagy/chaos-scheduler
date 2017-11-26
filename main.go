package main

import (
	//"fmt"
	"github.com/acnagy/chaos-scheduler/random"
	"github.com/acnagy/chaos-scheduler/threads"
	//"github.com/acnagy/chaos-scheduler/weather"
	"flag"
	"log"
	"os"
)

var threadpoolSize = flag.Int("t", 10, "set size for threadpool")
var weatherMode = flag.Bool("w", true, "run weather scheduling")
var randomMode = flag.Bool("r", true, "run random scheduling")

func main() {

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
	log.Println("[main] Logs ready to go!")

	// Determine running modes
	flag.Parse()
	threadpool := make([]threads.Thread, *threadpoolSize)
	threadpool = threads.InitThreadpool(threadpool, *threadpoolSize)
	threads.LogThreadpool(threadpool, *threadpoolSize)
	log.Printf("[main] Modes: weather: %t, random: %t, threadpool size: %d",
		*weatherMode, *randomMode, *threadpoolSize)

	if *randomMode {
		done := make(chan bool, 1)
		go random.Run(threadpool, *threadpoolSize, done)
		<-done
	}

	log.Println("[main] - complete!")

}
