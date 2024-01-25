package ejercicios

import (
	"fmt"
	"strconv"
)

func Funcion(valor string) (int, string) {

	entero, error := strconv.Atoi(valor)
	fmt.Println("error", error)
	if entero >= 100 {
		return entero, "Mayor o igual que 100"
	} else {
		return entero, "Menor que 100"
	}
}
