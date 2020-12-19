package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	mapArg := make(map[string]uint64) //мэп - где будет храниться два аргумента и конечный результат
	mapArg["firstArg"] = 0
	mapArg["secondArg"] = 0
	mapArg["result"] = 0

	fmt.Print("Введите номер числа фибоначчи:\n")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {

		num, err := strconv.ParseUint(input.Text(), 10, 64)

		if err != nil { //проверка на числовое значение
			fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
			os.Exit(1)
		}

		fibonacchi := count(num, mapArg)
		fmt.Printf("Число фибоначчи: %d", fibonacchi)

	}

}

func count(index uint64, mapArg map[string]uint64) uint64 {

	if index == 0 { //если index = 0, выходим из рекурсии.
		return mapArg["result"]
	}

	if mapArg["result"] == 0 && mapArg["firstArg"] == 0 && mapArg["secondArg"] == 0 { //ловим первый проход.
		mapArg["secondArg"] = 1 //при первом проходе result = 0, firstArg = 0, secondArg = 1.
		return count(index-1, mapArg)
	}

	mapArg["result"] = mapArg["firstArg"] + mapArg["secondArg"] //логика работы, пример: result = firstArg(5) + secondArg(3) = 8
	mapArg["secondArg"] = mapArg["firstArg"]                    //меняем местами. secondArg(3) = firstArg(5) , => secondArg = 5
	mapArg["firstArg"] = mapArg["result"]                       //firstArg(5) = result(8), firstArg => 8. На следующем шаге: result = firstArg(8)+ secondArg(5)

	return count(index-1, mapArg)

}
