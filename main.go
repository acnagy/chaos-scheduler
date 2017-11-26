package main

import (
	//"fmt"
	"github.com/acnagy/chaos-scheduler/random"
	"github.com/acnagy/chaos-scheduler/threads"
	//"github.com/acnagy/chaos-scheduler/weather"
	"flag"
	//"log"
	//"os"
)

var threadpoolSize = flag.Int("t", 10, "set size for threadpool")
var weatherMode = flag.Bool("w", true, "run weather scheduling")
var randomMode = flag.Bool("r", true, "run random scheduling")

func main() {

	flag.Parse()
	threadpool := make([]threads.Thread, *threadpoolSize)
	threadpool = threads.InitThreadpool(threadpool, *threadpoolSize)
	threads.PrintThreadpool(threadpool, *threadpoolSize)

	if *randomMode {
		done := make(chan bool, 1)
		go random.Run(threadpool, *threadpoolSize, done)
		<-done
	}

}

/*l, err := os.OpenFile("dev/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
      log.Fatal(err)
  }
  defer l.Close()

  log.SetOutput(l)
  log.Println("Logs ready to go!")*/
