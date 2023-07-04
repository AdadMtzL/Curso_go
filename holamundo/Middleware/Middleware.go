package middleware

import "fmt"

func Sumar(a int, b int) int {
	return a + b
}

func Restar(a int, b int) int {
	return a - b
}

func Multiplicar(a int, b int) int {
	return a * b
}

func MiMidlleware() {
	fmt.Println("Inicio")
	result := operacionesMid(Sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMid(Restar)(4, 2)
	fmt.Println(result)
	result = operacionesMid(Multiplicar)(6, 8)
	fmt.Println(result)
}

//Misma definición que las demás funciones
func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(a int, b int) int {
		//Aqui podemos supervisar los valores de la función
		fmt.Println("El valor de a es ", a)
		fmt.Println("El valor de a es ", b)
		return f(a, b)
	}
}

/*
Middelware: Función intermedia que se ejecuta andes de las demás, lo que reciben tiene que se lo mismo que devuelve
			Es un interceptor con los mismos parametros que regresa el control a la funcipn original
Se maneja differente con api que con otras cosas

*/
