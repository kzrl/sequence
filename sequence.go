package  sequence

import (
	"sync"
)

var num int
var m sync.Mutex

// Reset resets the counter to zero.
func Reset() int {
	m.Lock()
	defer m.Unlock()
	num = 0
	return num
}


func Next() int {
	m.Lock()
	defer m.Unlock()
	num++
	return num
}
