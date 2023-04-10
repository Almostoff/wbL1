package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func (h *Human) Greetings() {
	fmt.Printf("Hello, my name is %s and i`m %v yo./", h.name, h.age)
}

func main() {
	a := Action{Human{"Nick", 22}}
	a.Greetings() //унаследовала метод структуры Human
}
