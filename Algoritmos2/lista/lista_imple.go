package lista

type listaEnlazada struct {
	primerNodo *Nodo
	ultimoNodo *Nodo
	largo      int
}

func (l *listaEnlazada) EstaVacia() bool {
	if l == nil || l.largo == 0 {
		return true
	}

	return false
}

func (l *listaEnlazada) InsertarPrimero(v any) {
	nuevoNodo := &Nodo{v, nil}

	if l.EstaVacia() {
		l.primerNodo = nuevoNodo
		l.ultimoNodo = nuevoNodo
	} else {
		nuevoNodo.Siguiente = l.primerNodo
		l.primerNodo = nuevoNodo
	}
	l.largo++
}

func (l *listaEnlazada) InsertarUltimo(v any) {
	nuevoNodo := &Nodo{v, nil}

	if l.EstaVacia() {
		l.primerNodo = nuevoNodo
		l.ultimoNodo = nuevoNodo
	} else {
		l.ultimoNodo.Siguiente = nuevoNodo
		l.ultimoNodo = nuevoNodo
	}
	l.largo++
}

func (l *listaEnlazada) BorrarPrimero() any {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	res := l.primerNodo.Dato

	l.primerNodo = l.primerNodo.Siguiente

	if l.largo == 1 {
		l.ultimoNodo = nil
	}

	l.largo--
	return res
}

func (l *listaEnlazada) VerPrimero() any {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primerNodo.Dato
}

func (l *listaEnlazada) VerUltimo() any {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimoNodo.Dato
}

func (l *listaEnlazada) Largo() int {
	return l.largo
}

// Iterador interno
func (l *listaEnlazada) Iterar(visitar func(any) bool) {
	for nodoActual := l.primerNodo; nodoActual != nil; nodoActual = nodoActual.Siguiente {
		if !visitar(nodoActual.Dato) {
			return
		}
	}
}

// Devuelve iterador externo
func (l *listaEnlazada) Iterador() IteradorLista {
	iterador := &iteradorLista{l.primerNodo, nil, l}
	return iterador

}

func CrearLista[T any]() Lista {
	return new(listaEnlazada)
}
