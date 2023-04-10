package main

import (
	"fmt"
)

func reverseStr(str string) string {
	rStr := []rune(str)

	for i := 0; i < len(rStr)/2; i++ {
		rStr[i], rStr[len(rStr)-i-1] = rStr[len(rStr)-i-1], rStr[i]
	}
	return string(rStr)
}

func main() {
	str := "главрыба"
	fmt.Println("result: ", reverseStr(str))
}
