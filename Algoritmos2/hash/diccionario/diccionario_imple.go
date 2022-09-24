package diccionario

import (
	"bytes"
	"encoding/gob"
	"hash/lista"
)

type diccionario[K comparable, V any] struct {
	listas   []lista.Lista
	cantidad int
	hashFn   func(K) int
}

type record struct {
	clave any
	valor any
}

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (d *diccionario[K, V]) Guardar(clave K, dato V) {
	if d.doResize() {
		d.resize()
	}
	posicion := d.getPos(clave)
	datoNodo := record{clave, dato}
	// Si no existe la lista
	if d.listas[posicion] == nil {
		lista := lista.CrearLista[lista.Nodo]()
		lista.InsertarPrimero(datoNodo)
		d.listas[posicion] = lista
	} else {
		//existe
		iterador := d.listas[posicion].Iterador()

		// recorro a ver si esta la clave
		for actual := iterador.VerActual(); actual != nil; actual = iterador.Siguiente() {

			//TODO ARREGLAR
			//existe
			if (*actual).(record).clave == clave {
				*actual = datoNodo
				return
			}
		}

		// no existe
		d.listas[posicion].InsertarUltimo(datoNodo)
	}
	//updateo cantidad
	d.cantidad++
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (d *diccionario[K, V]) Pertenece(clave K) bool {
	posicion := d.getPos(clave)
	if d.listas[posicion] == nil {
		return false
	}

	iterador := d.listas[posicion].Iterador()
	for reg := iterador.VerActual(); reg != nil; reg = iterador.Siguiente() {
		if (*reg).(record).clave == clave {
			return true
		}
	}

	return false
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
// 'La clave no pertenece al diccionario'
func (d *diccionario[K, V]) Obtener(clave K) V {
	posicion := d.getPos(clave)
	if d.listas[posicion] == nil {
		panic("La clave no pertenece al diccionario")
	}

	iterador := d.listas[posicion].Iterador()
	for reg := iterador.VerActual(); reg != nil; reg = iterador.Siguiente() {
		if (*reg).(record).clave == clave {
			return (*reg).(record).valor.(V)
		}
	}

	panic("La clave no pertenece al diccionario")

}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (d *diccionario[K, V]) Borrar(clave K) (v V) {
	posicion := d.getPos(clave)
	if d.listas[posicion] == nil {
		panic("La clave no pertenece al diccionario")
	}

	iterador := d.listas[posicion].Iterador()
	for reg := iterador.VerActual(); reg != nil; reg = iterador.Siguiente() {
		if (*reg).(record).clave == clave {
			res := iterador.Borrar()
			d.cantidad--
			if d.listas[posicion].EstaVacia() {
				d.listas[posicion] = nil
			}
			return res.(record).valor.(V)
		}
	}

	panic("La clave no pertenece al diccionario")
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (d *diccionario[K, V]) Cantidad() int { return d.cantidad }

// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del
// mismo
func (d *diccionario[K, V]) Iterar(visitar func(clave K, dato V) bool) {

	// Itero listas
	for i := 0; i < len(d.listas); i++ {

		listaActual := d.listas[i]
		if listaActual == nil {
			continue
		}
		iterador := listaActual.Iterador()
		for reg := iterador.VerActual(); reg != nil; reg = iterador.Siguiente() {
			if !visitar((*reg).(record).clave.(K), (*reg).(record).valor.(V)) {
				return
			}
		}

	}
}

// Iterador devuelve un IterDiccionario para este Diccionario
func (d *diccionario[K, V]) Iterador() IterDiccionario[K, V] {
	size := len(d.listas) - 1
	if d.cantidad != 0 {
		for i := 0; i <= size; i++ {
			if d.listas[i] != nil {
				return &iterador_diccionario[K, V]{d, &d.listas[i], i, d.listas[i].Iterador()}
			}
		}
	}
	return &iterador_diccionario[K, V]{d, nil, size, nil}
}

func CrearHash[K comparable, T any](hash func(K) int) *diccionario[K, T] {
	// &nodo[T]{t, nil, ""}
	sliceListas := make([]lista.Lista, 13, 13)
	return &diccionario[K, T]{sliceListas, 0, hash}
}

func (d *diccionario[K, V]) getPos(clave K) int {
	return d.hashFn(clave) % cap(d.listas)

}

func convertirABytes[K comparable](clave K) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(clave)
	return buf.Bytes()
}

func (d *diccionario[K, V]) doResize() bool {
	return d.Cantidad() > cap(d.listas)*3
}

func (d *diccionario[K, V]) resize() {
	sliceListas := make([]lista.Lista, d.cantidad*10, d.cantidad*10)
	newDicc := &diccionario[K, V]{sliceListas, 0, d.hashFn}

	oldIterador := d.Iterador()
	for {
		key, value := oldIterador.VerActual()
		newDicc.Guardar(key, value)
		oldIterador.Siguiente()
		if !oldIterador.HaySiguiente() {
			break
		}
	}
	*d = *newDicc
}
