package ejer_interfaces

import (
	"fmt"
	"golang/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Println("Soy un humano y estoy respirando", hu.Sexo())
}
