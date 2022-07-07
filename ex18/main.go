package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	count uint
}

func NewCounter() *Counter {
	return &Counter{count: 0}
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) GetCount() uint {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

func main() {
	n := 9999

	wg := sync.WaitGroup{}
	wg.Add(n)

	counter := NewCounter()
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.GetCount())
}
