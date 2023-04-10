package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(i int, ch <-chan int) {
	for d := range ch {
		fmt.Printf("from worker %d listen this-> %v\n", i, d)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var nWorker int
	fmt.Println("enter workers count")
	fmt.Scan(&nWorker)

	ch := make(chan int)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < nWorker; i++ {
		go worker(i+1, ch)
	}

	for {
		select {
		case <-done:
			close(ch)
			fmt.Println("--graceful shutdown--")
			return
		default:
			ch <- rand.Intn(49)
		}
	}
}
