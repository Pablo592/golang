package defer_panic

import "fmt"

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es un mensaje que se ejecuta al final")
	fmt.Println("Este es el segundo mensaje")
}

func VemosPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			fmt.Println("Recuperandome del panic", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
