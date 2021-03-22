package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Args is count fibonachi
type Args struct {
	index     uint64
	firstArg  uint64
	secondArg uint64
	result    uint64
}

func main() {

	fmt.Print("Введите номер числа фибоначчи:\n")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {

		num, err := strconv.ParseUint(input.Text(), 10, 64)

		if err != nil { //проверка на числовое значение
			fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
			os.Exit(1)
		}

		args := &Args{index: num}
		fibonacchi := count(args)

		fmt.Printf("Число фибоначчи: %d", fibonacchi)
	}

}

func count(args *Args) uint64 {

	if args.index == 0 { //если index = 0, выходим из рекурсии.
		return args.result
	}

	args.index--

	if args.result == 0 && args.secondArg == 0 { //ловим первый проход.
		args.secondArg = 1
		return count(args)
	}

	args.result = args.firstArg + args.secondArg //логика работы, пример: result = firstArg(5) + secondArg(3) = 8
	args.secondArg = args.firstArg               //меняем местами. secondArg(3) = firstArg(5) , => secondArg = 5
	args.firstArg = args.result                  //firstArg(5) = result(8), firstArg => 8. На следующем шаге: result = firstArg(8)+ secondArg(5)

	return count(args)

}
