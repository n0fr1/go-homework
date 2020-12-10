package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Введите число, до которого будет проверка на простые числа: ")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {

		num, err := strconv.Atoi(input.Text())

		if err != nil {
			fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
			os.Exit(1)
		}

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
				fmt.Printf("%d\n", i)
			}
		}
	}

}
