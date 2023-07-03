package arreglosslices

import "fmt"

//Para asignar valores solo se pone = [10]{n1,n2,n3,n4,...} esto es para asignar con código duro
var tabla [10]int

func MostrarArreglo() {
	fmt.Println("Arreglo de entereos")
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	//Forma 2 de generar un arreglo tabla2:= [10]String{"s1","s2",...}
	//Matriz := [n][m]datatype
}

/**
*Slices
**/
var tb []int = []int{1, 3, 4, 2, 54, 2} //No se define el tamaño del arreglo
func Slcs() {
	fmt.Println(tb)
	//El slice ajusta las dimensiones del arreglos de manera automatica
}
