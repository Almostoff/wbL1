package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func uniq(str string) bool {
	charCount := make(map[rune]int)

	str = strings.ToLower(str)

	for _, v := range str {
		if charCount[v] > 0 {
			return false
		}
		charCount[v]++
	}
	return true
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("result: ", uniq(str))
}