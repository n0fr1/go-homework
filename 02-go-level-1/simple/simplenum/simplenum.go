package simplenum

import (
	"math"
)

//GetNum is printing simple numbs
func GetNum(num int) []int {

	var slice []int

	for i := 2; i <= num; i++ { // 1 и 0 - не являются простыми числами, поэтому их сразу исключаем.

		simNum := true //доп.флаг. Если истина - то число простое.

		getSqrt := math.Sqrt(float64(i))    //Вычисляем квадратный корень. Один из множителей - обязательно будет в интервале
		intSqrt := int(math.Round(getSqrt)) //строго до округленного квадратного корня.

		for k := 2; k <= intSqrt; k++ {
			if i%k == 0 {
				simNum = false
				break
			}
		}

		if simNum {
			slice = append(slice, i) //добавляем простые числа в слайс
			//fmt.Printf("%d\n", i)
		}
	}

	return slice
}
