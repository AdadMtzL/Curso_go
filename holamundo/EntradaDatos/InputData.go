package entradadatos

import (
	"bufio" //Estandar para endtrada de datos, archivos, etc.
	"fmt"
	"os"
	"strconv"
)

var n1 int
var n2 int
var cad string
var err error

func DeteccionEntrada() {
	fmt.Println("Indique n1")
	scanner := bufio.NewScanner(os.Stdin) //Lectura de datos por teclado
	if scanner.Scan() {
		n1, err = strconv.Atoi(scanner.Text()) //La entrada de bufio es texto
		//En caso de que el usuario se equivoque
		if err != nil {
			panic("El dato no es correcto" + err.Error()) //El panic termina el programa
		}
	}

	fmt.Println("Indique n2")
	if scanner.Scan() {
		n2, err = strconv.Atoi(scanner.Text()) //La entrada de bufio es texto
		//En caso de que el usuario se equivoque
		if err != nil {
			panic("El dato no es correcto" + err.Error()) //El panic termina el programa
		}
	}

	fmt.Println("Indique una cadena de texto")
	if scanner.Scan() {
		cad = scanner.Text()
	}

	fmt.Println(cad, n1*n2)

}
