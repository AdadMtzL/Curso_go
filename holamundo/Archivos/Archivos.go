package archivos

import (
	"fmt"
	"os"
	//"bufio"
	//"ioutil"
)

var texto string
var err error

func Arch() {
	fmt.Println("Manejo de archivos")

	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d * 4 = %d", i, i*4)
	}
	archivo, err := os.Create("./Archivos/tabla.txt")
	if err != nil {
		panic("Error al crear el archivo")
	}
	fmt.Println(archivo, texto)
	archivo.Close()
}
