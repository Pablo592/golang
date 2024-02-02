package middleware

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

func Restar(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}

func Dividir(a, b int) int {
	return a / b
}

func MiMiddleware() {
	fmt.Println("Ejecutando mi middleware")

	resultado := operacionesMidd(Sumar)(10, 5)
	fmt.Println("El resultado de la suma es", resultado)
	resultado = operacionesMidd(Restar)(10, 5)
	fmt.Println("El resultado de la resta es", resultado)
	resultado = operacionesMidd(Multiplicar)(10, 5)
	fmt.Println("El resultado de la multiplicacion es", resultado)
	resultado = operacionesMidd(Dividir)(10, 5)
	fmt.Println("El resultado de la division es", resultado)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de la operacion")
		return f(a, b)
	}
}
