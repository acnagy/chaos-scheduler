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

var threadpool_size = flag.Int("t", 10, "set size for threadpool")
var weather_mode = flag.Bool("w", true, "run weather scheduling")
var random_mode = flag.Bool("r", true, "run random scheduling")

func main() {

	flag.Parse()
	threadpool := make([]threads.Thread, *threadpool_size)
	threadpool = threads.Init_Threadpool(threadpool, *threadpool_size)
	threads.Print_Threadpool(threadpool, *threadpool_size)

	if *random_mode {
		done := make(chan bool, 1)
		go random.Run(threadpool, *threadpool_size, done)
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
