package Iteraciones

import (
	"fmt"
)

func CicloFor() {
	fmt.Println("Ejemplo de un ciclo for")
	/**
	for ifinito:
	for {

		break //Condición de término del ciclo
	}
	**/
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
