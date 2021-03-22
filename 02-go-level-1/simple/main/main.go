package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
<<<<<<< HEAD
	"testproject/simple/simplenum"
=======

	"github.com/n0fr1/go-homework/02-go-level-1/simple/simplenum"
>>>>>>> tests
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

<<<<<<< HEAD
		simplenum.GetSimpleNum(num)
=======
		result := simplenum.GetNum(num)
		fmt.Println(result)
>>>>>>> tests

	}

}
