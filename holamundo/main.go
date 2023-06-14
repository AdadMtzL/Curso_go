package main

//El package hae referencia a dentro de que carpeta estan los archivos go neceasrio
//por ejemplo para archivos go de variabes en una carpeta llamada variables e llamaria package go
//Si no esta dentro de una carpeta el package es main.

//que pedo, es con git ahuevo esto
import (
	"fmt"
	"Users/19286463/Desktop/Cursos/Desarrollodesde0/seccion8_Go/ejercicios/holamundo/Variables"
	//poner el path completo del proyecto
)
//Las funciones que se pueden acceder desde eafurea son Nombre y las internas son nombre
// Funcion principal de go, el main tiene que ir así ahuevo.
//El package es a huevo main en este caso, el nombre del archivo no.
func main() {
	fmt.Println("Primer mensaje")
	Variables.Numenteros()
}
