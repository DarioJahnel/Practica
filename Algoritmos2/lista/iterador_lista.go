package lista

type iteradorLista struct {
	nodoActual   *Nodo
	nodoAnterior *Nodo
	lista        *listaEnlazada
}

func (i *iteradorLista) VerActual() *any {
	if i.nodoActual == nil {
		panic("El iterador termino de iterar")
	}
	return &i.nodoActual.Dato
}

func (i *iteradorLista) HaySiguiente() bool {
	return i.nodoActual != nil
}
func (i *iteradorLista) Siguiente() *any {
	if i.nodoActual == nil {
		panic("El iterador termino de iterar")
	}
	i.nodoAnterior = i.nodoActual
	i.nodoActual = i.nodoActual.Siguiente
	if i.nodoActual == nil {
		return nil
	}
	return &i.nodoActual.Dato
}

func (i *iteradorLista) Insertar(v any) {
	nuevoNodo := &Nodo{v, i.nodoActual}

	if !i.HaySiguiente() {
		i.lista.ultimoNodo = nuevoNodo
	}
	i.nodoActual = nuevoNodo
	i.nodoAnterior.Siguiente = i.nodoActual

	i.lista.largo++
}

func (i *iteradorLista) Borrar() any {
	if i.nodoActual == nil {
		panic("El iterador termino de iterar")
	}
	res := i.nodoActual.Dato

	if i.HaySiguiente() {
		i.nodoActual = i.nodoActual.Siguiente //muevo puntero al que sigue
		if i.nodoAnterior == nil {
			// Primer nodo
			i.lista.primerNodo = i.nodoActual
		} else {
			// Si es primer nodo no hago nada con anterior
			i.nodoAnterior.Siguiente = i.nodoActual //al puntero anterior.siguiente lo redirijo al siguiente
		}
	} else {
		// Ultimo nodo
		i.nodoActual = nil // termino de iterar
		if i.nodoAnterior != nil {
			i.nodoAnterior.Siguiente = nil      // redirijo el anterior a nil
			i.lista.ultimoNodo = i.nodoAnterior // apunto ultimo nodo de la lista al anterior (o nil si fuera el unico elemento)
		} else {
			i.lista.primerNodo = nil
			i.lista.ultimoNodo = nil
		}
	}
	i.lista.largo--

	return res
}
