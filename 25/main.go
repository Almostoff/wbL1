package main

import (
	"fmt"
	"time"
)

func mySleep(ms int) {
	<-time.After(time.Duration(ms) * time.Millisecond)
}

func main() {
	fmt.Println("start prog")
	mySleep(1000)
	fmt.Println("prog done")
}
