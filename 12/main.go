package main

import "fmt"

func main() {
	m := make(map[string]bool)

	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, v := range strs {
		m[v] = true
	}

	for v := range m {
		fmt.Println(v)
	}
}
