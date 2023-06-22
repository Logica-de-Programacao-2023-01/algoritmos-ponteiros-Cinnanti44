package main

import "fmt"

func updateValue(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	var value int
	n := 5
	updateValue(&value, n)
	fmt.Println("Valor atualizado:", value) // Valor atualizado: 15
}
