package main

import (
	"fmt"
	"golang/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(100)
	fmt.Println("Estado", estado)
	fmt.Println("Texto", texto)
}
