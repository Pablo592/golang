package main

import (
	"fmt"
	"golang/ejercicios"
	"golang/variables"
	"runtime"
)

func main() {
	estado, texto := variables.ConviertoaTexto(100)
	fmt.Println("Estado", estado)
	fmt.Println("Texto", texto)

	if os := runtime.GOOS; os == "windows" {
		fmt.Println("Sistema Operativo", os)
	} else if os == "linux" {
		fmt.Println("Sistema Operativo", os)
	} else {
		fmt.Println("Sistema Operativo no soportado")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Sistema Operativo", os)
	case "linux":
		fmt.Println("Sistema Operativo", os)
	default:
		fmt.Println("Sistema Operativo no soportado")
	}

	entero, texto := ejercicios.Funcion("100")
	fmt.Println("Entero", entero)
	fmt.Println("Texto", texto)
}