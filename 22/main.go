package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1) //инициализирует как биг инт
	b := big.NewInt(1)

	a.SetString("12848372824939822", 10)
	b.SetString("9373297292738283", 10)

	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Println("sum: ", sum.String())

	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Println("sub: ", sub.String())

	mul := big.NewInt(0)
	mul.Mul(a, b)
	fmt.Println("mul: ", mul.String())

	div := big.NewInt(0)
	div.Div(a, b)
	fmt.Println("div: ", div.String())

}
