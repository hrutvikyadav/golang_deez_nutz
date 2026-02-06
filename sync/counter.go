package syncz

import "sync"

type Counter struct {
	mu sync.Mutex
	count int
}

func (c *Counter) Inc() {
	// WARN: without mutual exclusion we have race condition
	c.mu.Lock()
	defer c.mu.Unlock()
	//
	c.count++
}

func (c *Counter) CountValue() int {
	return c.count
}
