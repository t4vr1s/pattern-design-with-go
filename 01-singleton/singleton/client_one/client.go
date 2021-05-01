package client_one

import (
	"github.com/t4vr1s/cursos/patrones_de_diseño/01-singleton/singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.SetIncrementAge()
}
