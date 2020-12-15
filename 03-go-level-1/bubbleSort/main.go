package main

import (
	"fmt"
)

func main() {

	sliceNum := []int{22, -1, 6, 9, 10, 0, 4, 1, -3, 8, 7, 13, 48}

	sortedSlice := bubbleSort(sliceNum)

	fmt.Println("неотсортированный слайс: ")
	fmt.Println(sliceNum)

	fmt.Println("отсортированный слайс: ")
	fmt.Println(sortedSlice)

}

func bubbleSort(sliceNum []int) []int { //сортировка "пузырьком"

	var sliceSortNum = make([]int, len(sliceNum))
	copy(sliceSortNum, sliceNum)

	for j := 1; j < len(sliceSortNum); j++ {
		for el := 0; el < len(sliceSortNum)-j; el++ {

			if sliceSortNum[el] > sliceSortNum[el+1] {
				sliceSortNum[el], sliceSortNum[el+1] = sliceSortNum[el+1], sliceSortNum[el]
			}
		}

	}

	return sliceSortNum
}
