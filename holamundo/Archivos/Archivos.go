package archivos

import (
	"fmt"
)

var texto string

func Arch() string {
	fmt.Println("Manejo de archivos")

	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d * 4 = %d", i, i*4)
	}
	return texto
}
