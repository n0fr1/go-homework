package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var operation string

	fmt.Print("Введите первое число: ")
	a, err := testNumber() //проверяем 1-й введенный аргумент на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	b, err := testNumber() //проверяем 2-й введенный аргумент на число
	if err != nil {
		fmt.Fprint(os.Stderr, "Вы ввели не число! \n")
		os.Exit(1)
	}

	fmt.Print("Введите операцию: ")
	fmt.Scanln(&operation)

	if !(b == 0 && operation == "/") {
		view(a, b, operation)
	} else {
		fmt.Print("На ноль делить нельзя!")
	}

}

func view(a, b float64, calcOp string) {

	result, err := count(a, b, calcOp) //проверка на ошибку ввода знака операции
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	//здесь немного "поигрался" с выводом.
	if a == math.Round(a) && b == math.Round(b) { //если оба числа "равны" округленным выражениям
		fmt.Printf("%.0f %s %.0f %s %.0f", a, calcOp, b, "=", result) //то тогда нет смысла показывать нули после "запятой"
	} else if a == math.Round(+a) && b != math.Round(b) {
		fmt.Printf("%.0f %s %f %s %f", a, calcOp, b, "=", result)
	} else {
		fmt.Printf("%f %s %f %s %f", a, calcOp, b, "=", result)
	}
}

func count(a, b float64, calcOp string) (float64, error) {

	switch calcOp {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("Не выбрана операция из доступных: * / + - ")
	}
}

func testNumber() (float64, error) {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	num, err := strconv.ParseFloat(input.Text(), 64)
	return num, err

}
