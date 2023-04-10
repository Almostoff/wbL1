package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	m := sync.Map{} //потокобезопасная мапа
	//состоиз из двух мап - одна на запись, другая на чтение

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			m.Store(x, fmt.Sprintf("value %d", x))
			wg.Done()
		}(i)
	}

	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %d, value: %v\n", key, value)
		return true
	})
}
