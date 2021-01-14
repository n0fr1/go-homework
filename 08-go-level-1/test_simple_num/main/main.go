package main

import (
	"bufio"
	"fmt"
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

		simplenum.GetSimpleNum(num)

	}

}
