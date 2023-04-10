package main

import "fmt"

func intersect(a, b []string) []string {
	set1 := make(map[string]bool)
	for _, s := range a {
		set1[s] = true
	}

	result := []string{}
	for _, s := range b {
		if set1[s] {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	arr1 := []string{"apple", "banana", "orange", "pear"}
	arr2 := []string{"orange", "pear", "kiwi", "pineapple"}

	result := intersect(arr1, arr2)
	fmt.Println(result)
}
