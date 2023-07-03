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
	fmt.Fprintln(archivo, textoF)
	archivo.Close()
	//Para abrir un archivo esarch,err := os.OpenFile(FileName,Os.O_WRONLY | os.O_APPEND,permisos tipolinux) Esto es para cuando un archivo pueda seguir recibiendo texto
	//Escritura de texto tambien se puede  con arch.WriteString(texto)
	//Siempre cerrar el archivo
	//para leer archivos es os.Open(filename)
	// scanner := bufio.newscanner(archivo), luego imprimir mediante for y cerrar el archivo
	//for scanner.scan { letras := scanner.text print letras}
}
