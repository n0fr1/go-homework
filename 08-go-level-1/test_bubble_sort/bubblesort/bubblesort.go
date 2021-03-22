package bubblesort

//Bubble is slice-sorting function
func Bubble(sliceNum []int) []int { //сортировка "пузырьком"

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
