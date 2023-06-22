package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func preencherPrimos(s *[]int, tamanho int) {
	count := 0
	num := 2

	for count < tamanho {
		if isPrime(num) {
			*s = append(*s, num)
			count++
		}
		num++
	}
}

func main() {
	var primos []int
	tamanho := 10

	preencherPrimos(&primos, tamanho)

	fmt.Println("Números primos:", primos)
}
