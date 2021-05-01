package client_two

import (
	"github.com/t4vr1s/cursos/patrones_de_diseño/01-singleton/singleton"
)

func AgeIncrement() {
	p := singleton.GetInstance()
	p.SetIncrementAge()
}
