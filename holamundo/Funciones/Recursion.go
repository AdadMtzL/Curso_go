package funciones

import "fmt"

func Rec(valor int) {
	if valor > 9000 {
		return
	} else {
		fmt.Println(valor)
		Rec(valor * 2)
	}
}
