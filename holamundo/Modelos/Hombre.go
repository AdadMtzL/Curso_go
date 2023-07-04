package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
	Vivo       bool
}

func (H *Hombre) Respirar() {
	H.Respirando = true
}

func (H *Hombre) Comer() {
	H.Comiendo = true
}

func (H *Hombre) Pensar() {
	H.Pensando = true
}

func (H *Hombre) Genero() string {
	return "Hombre"
}
