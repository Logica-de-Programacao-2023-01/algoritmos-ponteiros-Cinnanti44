package main

import "fmt"

func sumLastTwoDigits(ptr *int) {
	num := *ptr                        // Obtém o valor da variável inteira através do ponteiro
	lastDigit := num % 10              // Obtém o último dígito
	secondLastDigit := num / 10 % 10   // Obtém o segundo último dígito
	sum := lastDigit + secondLastDigit // Calcula a soma dos dois últimos dígitos
	*ptr = sum                         // Atualiza o valor da variável com a soma
}

func main() {
	num := 1234
	sumLastTwoDigits(&num)
	fmt.Println("Novo valor:", num) // Novo valor: 7
}
