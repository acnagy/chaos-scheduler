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
var weatherStaticMode = flag.Bool("w", true, "run weather static scheduling")
var weatherVariableMode = flag.Bool("v", true, "run weather variable scheduling")
var sjfMode = flag.Bool("j", true, "run shortest-job-first scheduling")

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

	// Determine/Notify running modes
	flag.Parse()
	runConfig := fmt.Sprintf(
		"[main] modes: random: %t, weather - static: %t, weather - variable: %t, sjf: %t; threadpool size: %d, max threads: %d\n",
		*randomMode, *weatherStaticMode, *weatherVariableMode, *sjfMode, *threadpoolSize, *numberThreads,
	)
	log.Printf(runConfig)
	fmt.Printf(runConfig)

	// Setup running modes
	randomDone := make(chan bool, 1)
	randomThr := make(chan threads.Thread, *numberThreads)

	weatherStaticDone := make(chan bool, 1)
	weatherStaticThr := make(chan threads.Thread, *numberThreads)

	weatherVariableDone := make(chan bool, 1)
	weatherVariableThr := make(chan threads.Thread, *numberThreads)

	sjfDone := make(chan bool, 1)
	sjfThr := make(chan threads.Thread, *numberThreads)
	fmt.Println("make channels")

	// Init thread creation routine
	threadsDone := make(chan bool, 1)
	threads.InitWaitingThreads(randomThr, weatherStaticThr, weatherVariableThr, sjfThr, *threadpoolSize, *numberThreads)
	go threads.CreateThreadRandomly(
		randomThr, weatherStaticThr, weatherVariableThr, sjfThr,
		*numberThreads-*threadpoolSize,
		threadsDone,
	)

	// Run the selected mode(s)
	if *randomMode {
		go scheduling.Run("random", *threadpoolSize, *numberThreads, randomThr, randomDone)
	} else {
		randomDone <- true
	}

	if *weatherStaticMode {
		go scheduling.Run("weather - static", *threadpoolSize, *numberThreads, weatherStaticThr, weatherStaticDone)
	} else {
		weatherStaticDone <- true
	}

	if *weatherVariableMode {
		go scheduling.Run("weather - variable",
			*threadpoolSize,
			*numberThreads,
			weatherVariableThr,
			weatherVariableDone)
	} else {
		weatherVariableDone <- true
	}

	if *sjfMode {
		go scheduling.Run("sjf", *threadpoolSize, *numberThreads, sjfThr, sjfDone)
	} else {
		sjfDone <- true
	}

	<-randomDone
	<-weatherStaticDone
	<-weatherVariableDone
	<-sjfDone
	<-threadsDone

	log.Println("[main] - complete!")
}
