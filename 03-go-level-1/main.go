package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Printf("FizzBuzz %d\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Fizz %d\n", i)
		} else if i%5 == 0 {
			fmt.Printf("Buzz %d\n", i)
		}
	}

}
