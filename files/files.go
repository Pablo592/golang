package files

import (
	"bufio"
	"fmt"
	"golang/teclado"
	"os"
)

func GrabaTabla() {
	texto := teclado.TablaMultiplicar()
	filename := "files/tabla.txt"
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Print("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {

	var texto string = teclado.TablaMultiplicar()
	filename := "files/tabla.txt"
	if !append(filename, texto) {
		fmt.Println("Error al concatenar texto")
	}
}

func append(filename string, texto string) bool {
	archivo, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = archivo.WriteString(texto)

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	archivo.Close()
	return true
}

func LeerArchivo() {
	archivo, err := os.Open("files/tabla.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Hubo un error")
	}
}
