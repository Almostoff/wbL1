package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, target int) int {
	var min int
	max := len(nums) - 1

	for min <= max {
		mid := (max + min) / 2 // min + ((max - min) / 2) to fix
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			max = mid - 1
		case nums[mid] < target:
			min = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 4, 6, 7, 8, 11}
	findT := 3
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= findT
	})
	if index < len(arr) && arr[index] == findT {
		fmt.Printf("i`m found this// index is:  %d\n", index)
	} else {
		fmt.Printf("i`m not found for this check: %d\n", findT)
	}
	fmt.Println("From handFunc ", binarySearch(arr, findT))
	fmt.Println("From include method: ", sort.SearchInts(arr, findT)) //возвращает предположительный индекс где бы он мог стоять
}
