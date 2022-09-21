package lista

type Lista interface {
	EstaVacia() bool
	InsertarPrimero(any)
	InsertarUltimo(any)
	BorrarPrimero() any
	VerPrimero() any
	VerUltimo() any
	Largo() int
	Iterar(visitar func(any) bool)
	Iterador() IteradorLista
}

type IteradorLista interface {
	VerActual() *any
	HaySiguiente() bool
	Siguiente() *any
	Insertar(any)
	Borrar() any
}
