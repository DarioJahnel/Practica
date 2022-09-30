package diccionario

import (
	"bytes"
	"encoding/gob"
)

type Arbol[K comparable, V any] struct {
	raiz     *Nodo[K, V]
	cantidad int
}

type Nodo[K comparable, V any] struct {
	key      K
	value    V
	hijo_izq *Nodo[K, V]
	hijo_der *Nodo[K, V]
}

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (a *Arbol[K, V]) Guardar(clave K, dato V) {
	if a.guardarNodo(a.raiz, clave, dato) == nil {
		a.raiz = &Nodo[K, V]{clave, dato, nil, nil}
		a.cantidad++
	}
}

func (a *Arbol[K, V]) guardarNodo(raiz *Nodo[K, V], clave K, dato V) *Nodo[K, V] {
	if raiz == nil {
		return nil
	}
	if raiz.key == clave {
		raiz.value = dato
		return raiz
	}

	var res *Nodo[K, V]
	if comparar[K](raiz.key, clave) > 1 {
		res = a.buscarNodo(raiz.hijo_izq, clave)
		if res == nil {
			raiz.hijo_izq = &Nodo[K, V]{clave, dato, nil, nil}
			a.cantidad++
			return raiz.hijo_izq
		}
	} else {
		res = a.buscarNodo(raiz.hijo_der, clave)
		if res == nil {
			raiz.hijo_der = &Nodo[K, V]{clave, dato, nil, nil}
			a.cantidad++
			return raiz.hijo_der
		}
	}

	return res
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (a *Arbol[K, V]) Pertenece(clave K) bool {
	res := a.buscarNodo(a.raiz, clave)
	return res != nil
}

func (a *Arbol[K, V]) buscarNodo(raiz *Nodo[K, V], clave K) *Nodo[K, V] {
	if raiz == nil {
		return nil
	}
	if raiz.key == clave {
		return raiz
	}

	var res *Nodo[K, V]
	if comparar[K](raiz.key, clave) > 1 {
		res = a.buscarNodo(raiz.hijo_izq, clave)
	} else {
		res = a.buscarNodo(raiz.hijo_der, clave)
	}

	return res
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
// 'La clave no pertenece al diccionario'
func (a *Arbol[K, V]) Obtener(clave K) V {
	res := a.buscarNodo(a.raiz, clave)
	if res == nil {
		panic("La clave no pertenece al diccionario")
	}

	return res.value
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (a *Arbol[K, V]) Borrar(clave K) V {
	return a.borrarNodo(a.raiz, clave)
}

func (a *Arbol[K, V]) borrarNodo(raiz *Nodo[K, V], clave K) V {
	if raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	if raiz.key == clave {
		valor := raiz.value
		// Si no tiene hijos
		if raiz.hijo_izq == nil && raiz.hijo_der == nil {
			*raiz = Nodo[K, V]{}
		} else if raiz.hijo_izq != nil && raiz.hijo_der == nil {
			// Si tiene hijo izq pero no der
			raiz = raiz.hijo_izq

		} else if raiz.hijo_izq == nil && raiz.hijo_der != nil {
			// Si tiene hijo der pero no izq
			raiz = raiz.hijo_der

		} else if raiz.hijo_izq != nil && raiz.hijo_der != nil {
			// Si tiene 2 hijos
			reemplazo := a.encontrarReemplazo(raiz.hijo_izq)
			a.Borrar(reemplazo.key)
			raiz.key = reemplazo.key
			raiz.value = reemplazo.value
		}
		return valor
	}

	var res V
	if comparar[K](raiz.key, clave) > 1 {
		res = a.borrarNodo(raiz.hijo_izq, clave)
	} else {
		res = a.borrarNodo(raiz.hijo_der, clave)
	}

	return res
}

func (a *Arbol[K, V]) encontrarReemplazo(raiz *Nodo[K, V]) *Nodo[K, V] {
	if raiz == nil {
		return nil
	}

	res := a.encontrarReemplazo(raiz.hijo_der)
	if res == nil {
		// Encontre el que tengo que borrar y devolver
		return raiz
	}

	return res
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (a *Arbol[K, V]) Cantidad() int { return a.cantidad }

// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del
// mismo
func (a *Arbol[K, V]) Iterar(func(clave K, dato V) bool) {}

// Iterador devuelve un IterDiccionario para este Diccionario
func (a *Arbol[K, V]) Iterador() IterDiccionario[K, V] { return nil }

// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
// incluyéndolos en caso de encontrarse
func (a *Arbol[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {}

// IteradorRango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
func (a *Arbol[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] { return nil }

func convertirABytes[K comparable](clave K) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(clave)
	return buf.Bytes()
}

func comparar[K comparable](claveNodo K, claveInput K) int {
	nodoClave := convertirABytes[K](claveNodo)
	inputClave := convertirABytes[K](claveInput)

	return bytes.Compare(nodoClave, inputClave)

}

func CrearDiccionarioOrdenado[K comparable, V any]() Diccionario[K, V] {
	return &Arbol[K, V]{nil, 0}
}
