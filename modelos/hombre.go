package modelos

type Hombre struct {
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar()      { h.respirando = true }
func (h *Hombre) Comer()         { h.comiendo = true }
func (h *Hombre) Pensar()        { h.pensando = true }
func (h *Hombre) Sexo() string   { return "Hombre" }
func (h *Hombre) EstaVivo() bool { return h.vivo }
