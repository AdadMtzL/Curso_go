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
	for i < 10 { //i:=0;i<10;i++ es lo mismo i+=n incrementa en n
		//Se puede combinar con multiples estructuras como las
		//Condicionales asignación de variables y cosa así
		//hay instrucciones como break, continue, entre otras
		fmt.Println(i)
		i++
	}
}
