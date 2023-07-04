package main

//El package hae referencia a dentro de que carpeta estan los archivos go neceasrio
//por ejemplo para archivos go de variabes en una carpeta llamada variables e llamaria package go
//Si no esta dentro de una carpeta el package es main.
//comando para iniciarlo en go es go mod init y la ruta del archivo
//que pedo, es con git ahuevo esto
import (
	middleware "Users/19286463/Desktop/Cursos/Desarrollodesde0/seccion8_Go/Curso_go/holamundo/Middleware"
	"fmt"
	//strconv es un paquete para conversion a string
	//poner el path completo del proyecto
	//paqueet para detectar cosas del sistema
)

// Las funciones que se pueden acceder desde eafurea son Nombre y las internas son nombre
// Funcion principal de go, el main tiene que ir así ahuevo.
// El package es a huevo main en este caso, el nombre del archivo no.
func main() {
	//var str string
	//var no int
	fmt.Println("Primer mensaje")

	//	Variables.Numenteros()
	//	Variables.Resta()
	//	fmt.Println("Resta con parametros: ", Variables.Restaparametros(55, 12))
	//	if os := runtime.GOOS; os == "Linux." { //obtencion del sistema operativo y su comparacion eficiente
	//		fmt.Println("Sisema Linux o ios, no se la verdad")
	//	} else {
	//		fmt.Println("Sisema windows")
	//	}
	//para ponder anidas poner el if despues del else

	/**
	switch os := runtime.GOOS; os {
	case "Linux.":
		code
	case "opcion2":
		code
	default:
		ftm.printf("%s \n",os)
	}
	**/

	//no, str = Ejercicios.Ejer("500")
	//fmt.Println("Enviado: ", no)
	//fmt.Println("Recibido", str)
	//Variables.Resta()

	//entradadatos.DeteccionEntrada()

	//For infinito
	//for{
	// Instruction set for the loop goes here
	//	break
	//Sirve para reemplazar un while
	//}

	//Iteraciones.CicloFor()
	//archivos.Arch()

	//webserver.Webserver()

	middleware.MiMidlleware()
}

//git status git add . git commit git push
