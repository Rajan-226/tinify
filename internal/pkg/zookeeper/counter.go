package zookeeper

import "sync"

var (
	counter int64
	mu      sync.Mutex
)

func GetCounter() int64 {
	mu.Lock()
	defer mu.Unlock()

	counter++

	return counter
}
