package arreglosslices

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  44,
		"Chivas":       66,
		"Boca Juniosr": 30,
	}
	fmt.Println(campeonato)
	//Las claves se ordenan automaticamente
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, no. %d \n", equipo, puntaje)
	} //como un foreach

	delete(campeonato, "Real Madrid") //Eliminar cosas de un mapa
	fmt.Println(campeonato)

	//Comprobación de que existe algo en un mapa clave,bolleanvar:= mapname[clave] fmt.ptintf("clave %d eleemtno %t",clave,bolleanvar)
}
