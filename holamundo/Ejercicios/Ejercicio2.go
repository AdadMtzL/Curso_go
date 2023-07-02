package Ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n1 int
var err error

func Ej2() {
	fmt.Println("Ejercicio número 2")
	fmt.Println("Indique un número")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		n1, err = strconv.Atoi(scanner.Text()) //La entrada de bufio es texto
		//En caso de que el usuario se equivoque
		if err != nil {
			fmt.Println("Error, se volvera a intentar recibir el valor")
			Ej2()
		}
	}
	for i := 1; i < 11; i++ {
		fmt.Println(n1, i, n1*i)
	}
}
