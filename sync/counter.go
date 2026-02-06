package syncz

import "sync"

type Counter struct {
	mu sync.Mutex
	count int
}

// An indication to consumer of API to use cons instead of initialising Counter themselves
// Avoids mistakes if someone tries to use by value instead of by pointer (recommended for mutexes)
func NewCounter() *Counter {
	return &Counter{}
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
