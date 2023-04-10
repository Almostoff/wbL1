package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	arr := strings.Split(str, " ")

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	res := strings.Join(arr, " ")
	fmt.Print("result is : ", res)
}
