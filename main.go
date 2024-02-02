package main

import (
	"fmt"
	"golang/goroutines"
)

func main() {
	/*
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

		teclado.IngresoNumeros()
	*/

	//teclado.TablaMultiplicar()
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeerArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//fmt.Println(funciones.Factorial(5))
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()

	//arreglos_slices.Capacidad()
	//mapas.MuestroMapa()
	//users.AltaUsuario()
	//Pedro := new(modelos.Hombre)
	//ejer_interfaces.HumanosRespirando(Pedro)
	//Maria := new(modelos.Mujer)
	//ejer_interfaces.HumanosRespirando(Maria)

	//defer_panic.VemosPanic()
	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Evangelina", canal1)

	estado := <-canal1
	fmt.Println(estado)
}
