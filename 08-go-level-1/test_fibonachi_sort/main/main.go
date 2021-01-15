package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/n0fr1/go-homework/08-go-level-1/test_fibonachi_sort/fibo"
)

func main() {

	fmt.Print("Введите номер числа фибоначчи:\n")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {

		num, err := strconv.ParseUint(input.Text(), 10, 64)

		if err != nil { //проверка на числовое значение
			fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
			os.Exit(1)
		}

		fibonacchi := fibo.Count(num)
		fmt.Printf("Число фибоначчи: %d", fibonacchi)

	}

}
