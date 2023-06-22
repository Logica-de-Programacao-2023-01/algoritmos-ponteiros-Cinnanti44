package main

import (
	"fmt"
	"math"
)

func updateWithPiAverage(ptr *float64) {
	*pi := math.Pi // Obtem o valor da constante Pi
	*ptr = (*ptr + *pi) / 2 // Calcula a média aritmética entre o valor atual e Pi
}

func main() {
	num := 3.14
	updateWithPiAverage(&num)
	fmt.Println("Novo valor:", num) // Novo valor: 3.772
}
