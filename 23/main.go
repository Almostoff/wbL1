package main

import "fmt"

func main() {
	i := 3
	arr := []int{2, 4, 5, 6, 10, 1}
	fmt.Println("before changes: ", arr)
	arr = append(arr[0:i], arr[i+1:]...)
	fmt.Println("after changes: ", arr)
}
