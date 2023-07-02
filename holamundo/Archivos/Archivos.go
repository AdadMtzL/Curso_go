package archivos

import (
	"fmt"
	"os"
	//"bufio"
	//"ioutil"
)

func datos() string {
	var texto string
	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d * 4 = %d", i, i*4)
	}
	return texto
}

func Arch() {
	fmt.Println("Manejo de archivos")
	var textoF string
	textoF = datos()
	archivo, err := os.Create("./Archivos/tabla.txt")
	if err != nil {
		panic("Error al crear el archivo")
	}
	fmt.Println(archivo, textoF)
	archivo.Close()
}
