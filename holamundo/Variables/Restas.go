package Variables

import (
	"fmt"
)

/**
Hay multiples tipos de dato, por lo que hasy desde bool, string, float32, fecha, etc.
Para fechas importar time y usar time.Time
**/

// var global
func Resta() {
	//var local
	num1 := int8(5)
	num2 := int8(2)
	num3 := num1 - num2
	fmt.Println("Resta: ", num3)
}

// parametros return
func Restaparametros(param1 int8, param2 int8) int8 {
	return param1 - param2
}
