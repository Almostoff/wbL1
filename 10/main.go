package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -31}

	// создаем карту для хранения значений температур в подмножествах
	tempGroups := make(map[int][]float64)

	// проходим по всем значениям температур
	for _, temp := range temperatures {
		// вычисляем номер группы с шагом в 10 градусов для текущего значения температуры
		group := int(math.Trunc(temp/10)) * 10

		// добавляем значение температуры в соответствующую группу
		tempGroups[group] = append(tempGroups[group], temp)
	}

	var groups []int
	for group := range tempGroups {
		groups = append(groups, group)
	}
	sort.Ints(groups)
	for _, group := range groups {
		fmt.Printf("%d: %v\n", group, tempGroups[group])
	}
}
