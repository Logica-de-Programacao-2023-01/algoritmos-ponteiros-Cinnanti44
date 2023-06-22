package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(l *Livro) {
	l.Preco = l.Preco * 0.9 // Aplica um desconto de 10% no preço do livro
}

func main() {
	livro := Livro{
		Titulo: "Aventuras de Sherlock Holmes",
		Autor:  "Arthur Conan Doyle",
		Preco:  50.0,
	}

	fmt.Println("Preço antes do desconto:", livro.Preco) // Preço antes do desconto: 50.0

	aplicarDesconto(&livro)

	fmt.Println("Preço depois do desconto:", livro.Preco) // Preço depois do desconto: 45.0
}
