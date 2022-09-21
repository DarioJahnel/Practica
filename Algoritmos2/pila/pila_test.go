package pila_test

import (
	TDAPila "pila"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	assert.True(t, pila.EstaVacia(), "La pila no esta vacia")
	assert.Panics(t, func() { pila.VerTope() }, "VerTope deberia entrar en panico si no tiene elementos")
	assert.Panics(t, func() { pila.Desapilar() }, "Desapilar deberia entrar en panico si no tiene elementos")
}

func TestApilar(t *testing.T) {
	pila := generarPila()
	assert.Equal(t, pila.VerTope(), 10, "Valor incorrecto")
}

func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 10000; i++ {
		pila.Apilar(i)
	}
	for i := 10000; i >= 0; i-- {
		assert.Equal(t, pila.VerTope(), i, "Valor incorrecto")
		assert.Equal(t, pila.Desapilar(), i, "Valor incorrecto")
	}
	assert.True(t, pila.EstaVacia(), "La pila no quedo vacia")
	assert.Panics(t, func() { pila.VerTope() }, "VerTope deberia entrar en panico si no tiene elementos")
	assert.Panics(t, func() { pila.Desapilar() }, "Desapilar deberia entrar en panico si no tiene elementos")
}

func TestInvertirPila(t *testing.T) {
	test1 := generarPila()
	test1.Apilar(20)
	test1.InvertirPila()
	// Verifico que la pila se haya invertido (numero par de capacidad)
	assert.Equal(t, test1.Desapilar(), 10)
	assert.Equal(t, test1.Desapilar(), 20)
	// Numero impar
	test2 := generarPila()
	test2.Apilar(20)
	test2.Apilar(30)
	test2.InvertirPila()
	assert.Equal(t, test2.Desapilar(), 10)
	assert.Equal(t, test2.Desapilar(), 20)
	assert.Equal(t, test2.Desapilar(), 30)
}

func generarPila() TDAPila.Pila[int] {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	return pila
}
