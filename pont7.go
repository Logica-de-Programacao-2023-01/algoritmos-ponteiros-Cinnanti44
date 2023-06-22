package main

import "fmt"

type Livro struct {
	Título string
	Autor  string
}

func atualizarTítulo(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Título = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Título: "Livro X",
		Autor:  "Anônimo",
	}

	fmt.Println("Título antes:", livro.Título) // Título antes: Livro X

	atualizarTítulo(&livro)

	fmt.Println("Título depois:", livro.Título) // Título depois: Desconhecido
}
