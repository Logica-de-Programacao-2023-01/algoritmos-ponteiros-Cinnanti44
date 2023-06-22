package main

import (
	"fmt"
)

func reverseString(ptr *string) {
	runes := []rune(*ptr)                                 // Converte a string em um slice de runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Inverte os elementos do slice
		runes[i], runes[j] = runes[j], runes[i]
	}
	*ptr = string(runes) // Atualiza o valor da string com o reverso
}

func main() {
	str := "Hello, World!"
	reverseString(&str)
	fmt.Println("String Reversa:", str)
}
