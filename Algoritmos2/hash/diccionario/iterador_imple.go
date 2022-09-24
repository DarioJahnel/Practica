package diccionario

import (
	"hash/lista"
)

type iterador_diccionario[K comparable, V any] struct {
	diccionario    *diccionario[K, V]
	elementoActual *lista.Lista
	indice         int
	iteradorLista  lista.IteradorLista
}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
// el iterador hay un elemento.
func (i *iterador_diccionario[K, V]) HaySiguiente() bool {
	if i.iteradorLista == nil {
		return i.existeSiguienteEnLista()
	} else {
		if i.iteradorLista.HaySiguiente() {
			return true
		} else {
			return i.existeSiguienteEnLista()
		}
	}
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (i *iterador_diccionario[K, V]) VerActual() (K, V) {
	if i.elementoActual == nil || (*i.elementoActual).EstaVacia() {
		panic("El iterador termino de iterar")
	}
	// Si va a ser el primer elemento de la lista
	var aux record

	if i.iteradorLista == nil {
		aux = (*i.elementoActual).VerPrimero().(record)
	} else {
		// Si estaba dentro de la lista enlazada
		aux = (*i.iteradorLista.VerActual()).(record)
	}
	return aux.clave.(K), aux.valor.(V)
}

// Siguiente si HaySiguiente, devuelve la clave actual (equivalente a VerActual, pero únicamente la clave), y
// además avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe entrar en pánico con
// mensaje 'El iterador termino de iterar'
func (i *iterador_diccionario[K, V]) Siguiente() K {
	if i.elementoActual == nil || (*i.elementoActual).EstaVacia() {
		panic("El iterador termino de iterar")
	}
	res := (*i.iteradorLista.VerActual()).(record).clave.(K)
	// itero
	if i.iteradorLista.Siguiente(); !i.iteradorLista.HaySiguiente() {
		// si no hay siguiente busco el siguiente elemento de la lista
		i.indice++
		i.buscarSiguienteEnLista()
	}

	return res
}

func (i *iterador_diccionario[K, V]) buscarSiguienteEnLista() {
	i.iteradorLista = nil
	i.elementoActual = nil
	for j := i.indice; j < len(i.diccionario.listas); j++ {
		i.indice = j
		elem := i.diccionario.listas[j]
		if elem == nil {
			continue
		} else {
			i.elementoActual = &elem
			i.iteradorLista = (*i.elementoActual).Iterador()
			return
		}
	}
}

func (i *iterador_diccionario[K, V]) existeSiguienteEnLista() bool {
	for j := i.indice + 1; j < len(i.diccionario.listas); j++ {
		elem := i.diccionario.listas[j]
		if elem == nil {
			continue
		} else {
			return true
		}
	}

	return false
}
