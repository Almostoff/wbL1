package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{}

	v = 2

	switch v.(type) {
	case int:
		fmt.Println("type is // int")
	case string:
		fmt.Println("type is // string")
	case bool:
		fmt.Println("type is // bool")
	case chan int:
		fmt.Println("type is // channel")
	default:
		fmt.Println("idk what is this///")
	}
	// доп проверка
	fmt.Println("full name of type:", reflect.TypeOf(v)) //лезет под капот интерфейса и смотрит че там
}
