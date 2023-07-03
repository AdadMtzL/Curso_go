package funciones

import "fmt"

func f1(valor int) func() int { //closure
	num := valor
	secuencia := 0
	return func() int {
		secuencia++
		return num * secuencia
	}
}

func Llamarf1() {
	numm := 5
	tabla := f1(numm)
	for i := 1; i <= numm; i++ {
		fmt.Println(tabla())
	}
}
