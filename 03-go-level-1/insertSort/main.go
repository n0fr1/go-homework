package main

import (
	"fmt"
)

func main() {

	sliceNum := []int{2, 5, 6, 9, 10, 0, 4, 1, 3, 8, 7}

	sortedSlice := sort(sliceNum)

	fmt.Println("неотсортированный слайс: ")
	fmt.Println(sliceNum)

	fmt.Println("отсортированный слайс: ")
	fmt.Println(sortedSlice)

}

func sort(sliceNum []int) []int { //сортировка "вставками"

	var sliceSortNum = make([]int, len(sliceNum))
	copy(sliceSortNum, sliceNum)

	for i := 1; i <= len(sliceSortNum)-1; i++ {

		x := sliceSortNum[i]
		j := i
		for j > 0 && sliceSortNum[j-1] > x {
			sliceSortNum[j] = sliceSortNum[j-1]
			j = j - 1
		}

		sliceSortNum[j] = x
	}

	return sliceSortNum
}
