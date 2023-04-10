package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(d chan struct{}) { // первый способ
		fmt.Println("gopher with chan started")
		close(d)
	}(done)
	<-done
	fmt.Println("gopher with chan stopped")

	go func() { //второй
		fmt.Println("gopher with wg started")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("gopher with wg stopped")

	go func(ctx context.Context) { // завершается передача контекста в горутину и она закрывается
		fmt.Println("gopher with ctx started")
	}(ctx)
	time.Sleep(1 * time.Millisecond) //cancel чаще нужно время что бы успеть выполнить отмену, поэтмоу ждём немного
	cancel()                         //отмена контекста
	fmt.Println("gopher with ctx stopped")
}
