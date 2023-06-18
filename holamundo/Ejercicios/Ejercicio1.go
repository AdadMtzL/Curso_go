package Ejercicios

import (
	"fmt"
	"strconv"
)

func Ejer(Val string) (int, string) {
	var msj string
	var num int
	var e error
	num, e = strconv.Atoi(Val)
	if e != nil {
		fmt.Println("Error")
		//code here
	}
	if num > 100 {
		msj = "El número es mayor a 100"
	} else {
		msj = "El número es menor a 100"
	}
	return num, msj
}
