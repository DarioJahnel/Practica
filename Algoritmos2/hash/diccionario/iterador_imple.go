package diccionario

import "hash/lista"

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
	if (*i.elementoActual) == nil || (*i.elementoActual).EstaVacia() {
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
	if (*i.elementoActual) == nil || (*i.elementoActual).EstaVacia() {
		panic("El iterador termino de iterar")
	}
	var res K

	if i.iteradorLista == nil {
		// lo creo
		i.iteradorLista = (*i.elementoActual).Iterador()
		return (*i.iteradorLista.VerActual()).(K)
	} else {
		// itero
		if i.iteradorLista.HaySiguiente() {
			i.iteradorLista.Siguiente()
			return (*i.iteradorLista.VerActual()).(K)
		} else {
			// si no hay siguiente busco el siguiente elemento de la lista
			i.buscarSiguienteEnLista()
			i.iteradorLista = (*i.elementoActual).Iterador()
			return (*i.iteradorLista.VerActual()).(K)
		}
	}
	// if i.iteradorLista == nil {
	// 	res = (*i.elementoActual).VerPrimero().(record).clave.(K)
	// 	// Si la lista contiene mas de un nodo hago un iterador y lo avanzo a la segunda posicion
	// 	if (*i.elementoActual).Largo() > 1 {
	// 		i.iteradorLista = (*i.elementoActual).Iterador()
	// 		i.iteradorLista.Siguiente()
	// 	} else {
	// 		// Si no busco el siguiente elemento en la lista
	// 		i.buscarSiguienteEnLista()
	// 	}

	// } else {
	// 	// Si el iterador de lista tiene mas valores, avanzo
	// 	res = (*i.iteradorLista.VerActual()).(record).clave.(K)
	// 	if i.iteradorLista.HaySiguiente() {
	// 		i.iteradorLista.Siguiente()
	// 	} else {
	// 		// Si no tiene busco el siguiente elemento en la lista
	// 		i.buscarSiguienteEnLista()
	// 	}

	// }
	return res
}

func (i *iterador_diccionario[K, V]) buscarSiguienteEnLista() {
	i.iteradorLista = nil
	for j := i.indice; j < len(i.diccionario.listas); j++ {
		elem := i.diccionario.listas[j]
		if elem == nil {
			continue
		} else {
			i.elementoActual = &elem
			i.indice = j
			break
		}
	}

	panic("El iterador termino de iterar")
}

func (i *iterador_diccionario[K, V]) existeSiguienteEnLista() bool {
	for j := i.indice; j < len(i.diccionario.listas); j++ {
		elem := i.diccionario.listas[j]
		if elem == nil {
			continue
		} else {
			return true
		}
	}

	return false
}
