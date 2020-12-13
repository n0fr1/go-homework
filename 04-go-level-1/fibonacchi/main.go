package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var result uint64 = 0 //итоговый результат

	mapArg := make(map[string]uint64) //мэп - где будет храниться два аргумента.
	mapArg["a"] = 0
	mapArg["b"] = 0

	fmt.Print("Введите номер числа фибоначчи:\n")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {

		num, err := strconv.ParseUint(input.Text(), 10, 64)

		if err != nil { //проверва на числовое значение
			fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
			os.Exit(1)
		}

		fibonacchi := count(num, result, mapArg)
		fmt.Printf("Число фибоначчи: %d", fibonacchi)

	}

}

func count(num, result uint64, mapArg map[string]uint64) uint64 {

	if num == 0 { //если num = 0, выходим из рекурсии.
		return result
	}

	if result == 0 && mapArg["a"] == 0 && mapArg["b"] == 0 { //вычисляем первый проход.
		mapArg["b"] = 1 //при первом проходе а = 0, b = 1.
		return count(num-1, 0, mapArg)
	}

	result = mapArg["a"] + mapArg["b"] //логика работы, пример: result = a(5) + b(3) = 8
	mapArg["b"] = mapArg["a"]          //меняем местами. b(3) = a(5) , => b = 5
	mapArg["a"] = result               //a(5) = result(8), a => 8. На следующем шаге: result = a(8)+ b(5)

	return count(num-1, result, mapArg)

}
