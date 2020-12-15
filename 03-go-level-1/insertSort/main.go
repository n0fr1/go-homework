package main

import (
	"fmt"
)

func main() {

	sliceNum := []int{2, 5, 6, -1, -2, 0, 4, 1, -3, 8, 7}

	sortedSlice := sort(sliceNum)

	fmt.Println("неотсортированный слайс: ")
	fmt.Println(sliceNum)

	fmt.Println("отсортированный слайс: ")
	fmt.Println(sortedSlice)

}

func sort(sliceNum []int) []int { //сортировка "вставками"

	var sliceSortNum = make([]int, len(sliceNum))
	copy(sliceSortNum, sliceNum)

	for i := 1; i < len(sliceSortNum); i++ {

		el := sliceSortNum[i]
		ind := i
		for ind > 0 && sliceSortNum[ind-1] > el {
			sliceSortNum[ind] = sliceSortNum[ind-1]
			ind--
		}

		sliceSortNum[ind] = el
	}

	return sliceSortNum
}
