package singleton

import (
	"sync"
)

var (
	p    *person
	once sync.Once
)

func GetInstance() *person {
	// se ejecuta una sola vez en todo el proyecto incluyendo goroutine
	once.Do(func() {
		p = &person{}
	})
	return p
}

// se crea privada para no crear instancias
type person struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *person) SetName(name string) {
	p.mux.Lock()
	p.name = name
	defer p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()

	return p.name
}

func (p *person) SetIncrementAge() {
	p.mux.Lock()
	p.age++
	defer p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
