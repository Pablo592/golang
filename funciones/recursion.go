package funciones

func Factorial(numero int) int {
	if numero == 1 {
		return 1
	}
	return numero * Factorial(numero-1)
}
