package main

import (
	"fmt"

	"github.com/n0fr1/go-homework/08-go-level-1/test_bubble_sort/bubblesort"
)

func main() {

	sliceNum := []int{22, -1, 6, 9, 10, 0, 4, 1, -3, 8, 7, 13, 48}

	sortedSlice := bubblesort.Bubble(sliceNum)

	fmt.Println("неотсортированный слайс: ")
	fmt.Println(sliceNum)

	fmt.Println("отсортированный слайс: ")
	fmt.Println(sortedSlice)

}
