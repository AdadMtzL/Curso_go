package rutinas

import (
	"fmt"
	"strings"
	"time"
)

// Una rutina se ejecuta en segundo plano
func SlowName(Name string) {
	letras := strings.Split(Name, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //milis
		fmt.Println(letra)
	}
	//Para llamar una subrutina solo pones go antes de la función, es para subprocesos por decirlo así
	//Finaliza cuando el proceso principal finaliza el programa
}

/**
*Canales son aquellos que permiten que rutinas iteractuen
*nombre chan tipodato
**/
func SlowName2(Name string, canal1 chan bool) {
	letras := strings.Split(Name, "")
	for _, letra := range letras {
		time.Sleep(500 * time.Millisecond) //milis
		fmt.Println(letra)
	}
	canal1 <- true
}

//Para canales se usa lo siguiente en el main:
//canal1 := make(chan bool)
//go func
//func 2
//<- canal1 //Similar a un awake o esperar que acabe la función principal
// Ó defer func() {<-canal1}
//Son asincronas las rutinas
