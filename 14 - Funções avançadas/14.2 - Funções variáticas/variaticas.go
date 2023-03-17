package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	total := soma(20, 30, 50, 100, 100)
	fmt.Println(total)
}
