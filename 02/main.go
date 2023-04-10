package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{2, 4, 6, 8, 10}

	wg.Add(len(arr))

	for _, v := range arr {
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(n * n)
		}(v, wg)
	}
	wg.Wait()
}
