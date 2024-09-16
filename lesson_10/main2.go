package main

import (
	"fmt"
	"sync"
	"time"
)

type CounterRW struct {
	mu sync.RWMutex
	c  map[string]int
}

func (c *CounterRW) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *CounterRW) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c
}

func main() {
	key := "test"

	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}
