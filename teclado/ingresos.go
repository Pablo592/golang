package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese numero 1")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2")
	scanner = bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda")
	scanner = bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()

		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}

	fmt.Println(leyenda, numero1*numero2)

}

func TablaMultiplicar() string {
	numero := 50
	texto := ""

	for i := 0; i < 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}
	return texto
}
