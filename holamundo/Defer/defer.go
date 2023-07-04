package Defer

import (
	"fmt"
	"log"
)

// Instrucción de go que permite determinar la ultima acción antes de salir de una función
func Example() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("Mensaje final")
	fmt.Println("Segundo mensaje")
}
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("fallo: %v", reco)
		}
	}()
	panic("adios")
}
