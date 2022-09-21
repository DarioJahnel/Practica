package cola_test

import (
	TDACola "cola-enlazada"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearCola[int]()
	assert.True(t, cola.EstaVacia(), "La cola deberia estar vacia")
	assert.Panics(t, func() { cola.VerPrimero() }, "Cola vacia, panic expected")
	assert.Panics(t, func() { cola.Desencolar() }, "Cola vacia, panic expected")

}

func TestEncolar(t *testing.T) {
	cola := TDACola.CrearCola[int]()
	fmt.Println(cola)
	cola.Encolar(10)
	assert.False(t, cola.EstaVacia())
	assert.Equal(t, cola.VerPrimero(), 10)
}

func TestCompleto(t *testing.T) {
	cola := TDACola.CrearCola[int]()
	for i := 0; i < 10000; i++ {
		cola.Encolar(10)
		cola.Encolar(20)
		assert.False(t, cola.EstaVacia())
		assert.Equal(t, cola.VerPrimero(), 10)
		assert.Equal(t, cola.Desencolar(), 10)
		assert.Equal(t, cola.VerPrimero(), 20)
		assert.Equal(t, cola.Desencolar(), 20)
		assert.True(t, cola.EstaVacia())
		assert.Panics(t, func() { cola.VerPrimero() }, "Cola vacia, panic expected")
		assert.Panics(t, func() { cola.Desencolar() }, "Cola vacia, panic expected")
	}
}
