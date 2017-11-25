package data

import (
	"math/rand"
	"time"
)

func Retrieve() (id uint16, priority uint16) {

	rand.Seed(time.Now().UnixNano())
	id = uint16(rand.Uint32())
	priority = uint16(rand.Uint32())

	return id, priority

}
