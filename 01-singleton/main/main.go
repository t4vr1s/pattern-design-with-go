package main

import (
	"fmt"
	"sync"

	"github.com/t4vr1s/cursos/patrones_de_diseño/01-singleton/singleton"

	"github.com/t4vr1s/cursos/patrones_de_diseño/01-singleton/singleton/client_one"
	"github.com/t4vr1s/cursos/patrones_de_diseño/01-singleton/singleton/client_two"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()

		go func() {
			defer wg.Done()
			client_two.AgeIncrement()
		}()
	}
	wg.Wait()
	p := singleton.GetInstance()
	fmt.Printf("p es: %T \n", p)
	age := p.GetAge()
	fmt.Println(age)
}
