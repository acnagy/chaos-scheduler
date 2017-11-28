package threads

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestsortThreads(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	thpool := InitThreadpoolControl()
	sort := func(thpool []Thread) []Thread {
		for i := 0; i < len(thpool); i++ {
			for j := 0; j < len(thpool)-1; j++ {
				if thpool[j].Priority > thpool[j+1].Priority {

					temp := thpool[j]
					thpool[j] = thpool[j+1]
					thpool[j+1] = temp
				}
			}
		}
		return thpool
	}

	thpoolSorted := sort(thpool)
	thpoolSortedTest := sortThreads(thpool)

	for i := 0; i < len(thpoolSortedTest); i++ {
		if thpoolSorted[i].Id != thpoolSortedTest[i].Id {
			t.Fail()
		}
	}
}

func TestWork(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	thpool := InitThreadpoolControl()

	var expectedWorktime int
	for i := 0; i < len(thpool); i++ {
		expectedWorktime = expectedWorktime + thpool[i].Worktime
	}
	_, duration := Work("testing", thpool)

	if duration != time.Duration(expectedWorktime) {
		fmt.Printf("duration: %d, expectedWorktime: %d", duration, expectedWorktime)
		t.Fail()
	}

}
