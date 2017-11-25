package main

import (
	"fmt"
	"github.com/acnagy/chaos-scheduler/random"
	//"github.com/acnagy/chaos-scheduler/weather"
	"flag"
	"log"
	"os"
)

type thread struct {
	id       uint16
	priority uint16
}

func main() {

	var threadpool_size = flag.Int("t", 10, "set size for threadpool")
	flag.Parse()
	threadpool := make([]thread, *threadpool_size)

	l, err := os.OpenFile("dev/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.SetOutput(l)
	log.Println("Logs ready to go!")

	fmt.Println("hello!")
	//weather.Retrieve()

	for i := 0; i < *threadpool_size; i++ {
		id, priority := data.Retrieve()
		th := thread{id, priority}

		fmt.Printf("id: %5d - priority: %5d\n", th.id, th.priority)
		threadpool[i] = th
	}

	fmt.Println(threadpool)

}
