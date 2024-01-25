package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Juan"
	Estado = true
	Sueldo = 10000.50
	Fecha = time.Now()
	fmt.Println("Nombre", Nombre)
	fmt.Println("Estado", Estado)
	fmt.Println("Sueldo", Sueldo)
	fmt.Println("Fecha", Fecha)

}

func ConviertoaTexto(numero int) (bool, string) {
	//var texto string				//texto := ""
	//texto = strconv.Itoa(numero) //strconv.Itoa(numero) convierte un entero a string
	texto := strconv.Itoa(numero)
	return true, texto

}
