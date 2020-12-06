package main

import (
	"fmt"
)

func main() {

	sliceNum := []int{2, 5, 6, 9, 10, 0, 4, 1, 3, 8, 7} //зачем нужен [...]

	fmt.Println("неотсортированный слайс: ")
	fmt.Println(sliceNum)

	sortedSlice := sort(sliceNum) //нужно ли использовать указатель на слайс?

	fmt.Println("отсортированный слайс: ")
	fmt.Println(sortedSlice)

}

func sort(sliceNum []int) []int { //сортировка "пузырьком"

	var sliceSortNum = make([]int, len(sliceNum))
	copy(sliceSortNum, sliceNum)

	for j := 1; j <= len(sliceSortNum)-1; j++ {

		for i := 0; i <= len(sliceSortNum)-1-j; i++ {

			if sliceSortNum[i] > sliceSortNum[i+1] {
				sliceSortNum[i], sliceSortNum[i+1] = sliceSortNum[i+1], sliceSortNum[i]
			}
		}

	}

	return sliceSortNum
}
