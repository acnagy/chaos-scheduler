package threads

import (
	"testing"
)

func TestBubbleSortThreads(t *testing.T) {
	thpool := []Thread{
		{6, 7, 8},
		{1, 2, 3},
		{5, 13, 14},
		{3, 4, 5},
		{2, 10, 11},
	}
	thpoolSorted := []Thread{
		{1, 2, 3},
		{3, 4, 5},
		{6, 7, 8},
		{2, 10, 11},
		{5, 13, 14},
	}
	thpoolSortedTest := bubbleSortThreads(thpool, len(thpool))

	for i := 0; i < len(thpoolSortedTest); i++ {
		if thpoolSorted[i].Id != thpoolSortedTest[i].Id {
			t.Fail()
		}
	}
}

func TestWork(t *testing.T) {
	thpool := []Thread{
		{6, 7, 8},
		{1, 2, 3},
		{5, 13, 14},
		{3, 4, 5},
		{2, 10, 11},
	}

	var expectedWorktime int
	for i := 0; i < len(thpool); i++ {
		expectedWorktime = expectedWorktime + thpool[i].Worktime
	}
	worktime := Work("testing", thpool, len(thpool))

	if worktime != expectedWorktime {
		t.Fail()
	}

}
