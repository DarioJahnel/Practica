package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaEnlazada[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaEnlazada[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaEnlazada[T]) VerTope() T {
	if p.cantidad == 0 {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaEnlazada[T]) Apilar(t T) {
	p.datos = append(p.datos, t)
	p.cantidad++
}

func (p *pilaEnlazada[T]) Desapilar() T {
	if p.cantidad == 0 {
		panic("La pila esta vacia, no se puede desapilar")
	}
	elem := p.datos[p.cantidad-1]
	p.cantidad--
	return elem
}

func CrearPilaDinamica[T any]() Pila[T] {
	return new(pilaEnlazada[T])
}

func (p *pilaEnlazada[T]) InvertirPila() {
	for i := 0; i < p.cantidad/2; i++ {
		posFinal := p.cantidad - 1 - i
		aux := p.datos[i]
		p.datos[i] = p.datos[posFinal]
		p.datos[posFinal] = aux
	}
}
