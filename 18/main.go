package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	v int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	c.v++
	c.Unlock()
}

func main() {
	c := &Counter{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 123; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("counter is: ", c.v)
}
