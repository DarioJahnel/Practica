package cola

type nodo[T any] struct {
	dato T
	next *nodo[T]
}

type colaEnlazada[T any] struct {
	primerNodo *nodo[T]
	ultimoNodo *nodo[T]
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primerNodo == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.primerNodo.dato
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (c *colaEnlazada[T]) Encolar(t T) {
	nuevoNodo := &nodo[T]{t, nil}
	if c.EstaVacia() {
		c.primerNodo = nuevoNodo
	} else {
		c.ultimoNodo.next = nuevoNodo
	}
	c.ultimoNodo = nuevoNodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia"
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	res := c.primerNodo.dato
	c.primerNodo = c.primerNodo.next
	return res
}

func CrearCola[T any]() Cola[T] {
	return new(colaEnlazada[T])
}
