package main

import "fmt"

func modificarValor(p *int) {
	*p = 42 // Modifica o valor da variável através do ponteiro
}

func main() {
	v := 10
	fmt.Println("Valor antes:", v) // Valor antes: 10

	modificarValor(&v)

	fmt.Println("Valor depois:", v) // Valor depois: 42

	// Libera a memória (opcional neste caso, pois a variável é do escopo da função main)
	// Caso o ponteiro aponte para uma área de memória alocada dinamicamente, é importante
	// liberar a memória quando não for mais necessária.
}
