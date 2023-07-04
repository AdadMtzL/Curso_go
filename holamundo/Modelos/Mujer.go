package modelos

type Mujer struct { //Hrencoa
	Hombre
}

func (M *Mujer) Respirar() {
	M.Respirando = true
}

func (M *Mujer) Comer() {
	M.Comiendo = true
}

func (M *Mujer) Pensar() {
	M.Pensando = true
}

func (M *Mujer) Genero() string {
	return "Mujer"
}
