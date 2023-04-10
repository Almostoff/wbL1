package main

import (
	"fmt"
	"time"
)

func after(d time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(d)
		ch <- true
	}()
	return ch
}

func main() { //просто не останавливается программа
	//var n int
	//fmt.Print("Введите время в секундах: ")
	//fmt.Scan(&n)

	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; ; i++ {
			c <- i
		}
	}()

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-after(1 * time.Millisecond):
			fmt.Println("Time is up!")
			return
		}
	}
}
