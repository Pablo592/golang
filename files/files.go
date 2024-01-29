package files

import (
	"fmt"
	"golang/teclado"
	"os"
)

func GrabaTabla() {
	texto := teclado.TablaMultiplicar()
	archivo, err := os.Create("files/tabla.txt")

	if err != nil {
		fmt.Print("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {

}
