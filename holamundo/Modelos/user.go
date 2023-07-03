package modelos

import "time"

type User struct { //Ejemplo de elementos de una clase
	Id        int
	Name      string
	CreatedAt time.Time
	status    bool
}

// Se complenenta con la carpeta modelos/Users

//Metodo de clase como el constructor, en vez de usuario pude ser cualquier cosa menos this o self
func (Usuario *User) AddUser(id int, name string, createdat time.Time, Status bool) {
	Usuario.Id = id
	Usuario.Name = name
	Usuario.CreatedAt = createdat
	Usuario.status = Status
}
