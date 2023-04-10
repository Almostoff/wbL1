package main

import "fmt"

func main() {
	var a, b = 5, 12
	fmt.Printf("before changes: %v %v\n", a, b)
	//a, b = b, a
	//fmt.Printf("after changes: %v %v\n", a, b)
	//a = a + b        актуально?
	//b = a - b
	//a = a - b
	//fmt.Printf("after changes: %v %v\n", a, b)
	//Это XOR (по сути тоже сложение)
	//a = a ^ b
	//b = a ^ b
	//a = a ^ b
	//fmt.Println("after changes: ", a, b)
}
