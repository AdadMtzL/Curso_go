package funciones

import "fmt"

//Función anonima
func An1() {
	//Las variables declaradas aqui antes de la función anonima pueden interactuar con esta
	//De manera simple
	suma := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(suma(5, 88))

	//Si se vuelve a declarar  la misma función suma es posible cambiar su funcionamiento, más
	//No la estructura
	//Puedes poner como retorno una función poniendo func nombre() func n2(params) return { }
}
