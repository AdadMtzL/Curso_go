package users

import (
	modelos "Users/19286463/Desktop/Cursos/Desarrollodesde0/seccion8_Go/Curso_go/holamundo/Modelos"
	"fmt"
	"time"
)

// Se complenenta con la carpeta modelos/Users
func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
