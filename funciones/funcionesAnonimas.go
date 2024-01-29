package funciones

import "fmt"

func Calculos() {
	suma := func(a int, b int) int {
		return a + b
	}
	fmt.Println(suma(5, 7))
}
