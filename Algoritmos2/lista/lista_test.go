package lista_test

import (
	TDALista "lista"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	assert.True(t, lista.EstaVacia(), "La lista no esta vacia")
	assert.Equal(t, lista.Largo(), 0)
	assert.Panics(t, func() { lista.BorrarPrimero() }, "BorrarPrimero deberia entrar en panico si no tiene elementos")
	assert.Panics(t, func() { lista.VerPrimero() }, "VerPrimero deberia entrar en panico si no tiene elementos")
	assert.Panics(t, func() { lista.VerUltimo() }, "VerUltimo deberia entrar en panico si no tiene elementos")
}

func TestUnicoElemento(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	lista.InsertarPrimero(10)
	assert.Equal(t, lista.VerPrimero(), 10)
	assert.Equal(t, lista.VerUltimo(), 10)
	assert.False(t, lista.EstaVacia())
	assert.Equal(t, lista.Largo(), 1)
	assert.Equal(t, lista.BorrarPrimero(), 10)
}

func TestDosElementos(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	assert.Equal(t, lista.VerPrimero(), 10)
	assert.Equal(t, lista.VerUltimo(), 20)
	assert.Equal(t, lista.BorrarPrimero(), 10)
	assert.Equal(t, lista.VerPrimero(), 20)
	assert.Equal(t, lista.VerUltimo(), 20)
	assert.Equal(t, lista.BorrarPrimero(), 20)
	assert.True(t, lista.EstaVacia())
}

func TestTresElementosIterador(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	iterador := lista.Iterador()
	assert.Equal(t, *iterador.Siguiente(), 20)
	assert.Equal(t, iterador.Borrar(), 20)
	assert.Equal(t, iterador.Borrar(), 30)
	assert.Panics(t, func() { iterador.Borrar() }, "El iterador deberia entrar en panico si intenta borrar y llego al final")
}

func TestTresElementosIteradorBorrar(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	iterador := lista.Iterador()
	assert.Equal(t, iterador.Borrar(), 10)
	assert.Equal(t, iterador.Borrar(), 20)
	assert.Equal(t, iterador.Borrar(), 30)
	assert.Panics(t, func() { iterador.Borrar() }, "El iterador deberia entrar en panico si intenta borrar y llego al final")
	assert.Panics(t, func() { lista.VerPrimero() }, "VerPrimero deberia entrar en panico si no tiene elementos")
}

func TestDosElementosIterador(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	iterador := lista.Iterador()
	assert.Equal(t, iterador.Borrar(), 10)
	assert.Equal(t, iterador.Borrar(), 20)
	assert.Panics(t, func() { iterador.Borrar() }, "El iterador deberia entrar en panico si intenta borrar y llego al final")
}

func TestIteradorVacio(t *testing.T) {
	lista := TDALista.CrearLista[int]()
	iterador := lista.Iterador()
	assert.Panics(t, func() { iterador.VerActual() }, "Iterador deberia arrojar panic")
	assert.Panics(t, func() { iterador.Siguiente() }, "Iterador deberia arrojar panic")
	assert.False(t, iterador.HaySiguiente())

}
