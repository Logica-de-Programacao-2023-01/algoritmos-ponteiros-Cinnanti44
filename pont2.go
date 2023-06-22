package main

import "fmt"

func updateValue(ptr *int) {
	if *ptr%2 == 0 { // verifica se o valor apontado pelo ponteiro é par
		*ptr = 0 // atualiza o valor para 0
	} else {
		*ptr = 1 // atualiza o valor para 1
	}
}

func main() {
	var value int
	value = 7
	updateValue(&value)
	fmt.Println("Valor atualizado:", value) // Valor atualizado: 1
}
