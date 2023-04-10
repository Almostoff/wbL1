package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// quicksort - sorts array with O(n*log(n))
func quicksort(a []int) []int { //как работает квиксорт
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	rand.NewSource(time.Now().UnixNano()) //новая версия рандома
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	a := []int{3, 7, 2, 9, 0, 1, 6, 8, 4}
	fmt.Println(a)
	fmt.Println("quicksort method out", quicksort(a))
	sort.Ints(a)
	fmt.Println("quicksort method in", a)
}
