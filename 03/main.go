package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := 0
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	for _, v := range arr {
		go func(a int) {
			res += a * a
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Printf("result eq %v", res)
}
