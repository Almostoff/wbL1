package main

import "fmt"

func setBit(n int64, i uint, b bool) {
	if b {
		n |= 1 << i // сдвигаю на i-ое кол-во бит и плюсую этот бит
		// для i = 3 будет 1 << 3 -> 100 и это плюсую к числу изначальному в битовом "пространстве"
	} else {
		n &^= 1 << i // сбрасываю бит в 0
	}
	fmt.Printf("число после изменения в двоичном виде: %b\n", n)
}

func main() {
	num := int64(41)
	fmt.Printf("число в двоичном виде: %b\n", num)
	setBit(num, 4, true)

	setBit(num, 3, false)
}
