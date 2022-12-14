package model

import (
	"admmemoria/administrador"
	"fmt"
)

type Persona struct {
	nombre    string
	hij_mayor *Persona
	hij_menor *Persona
	padre     *Persona
}

// CrearPersona devuelve un puntero a una nueva Persona, con el nombre y el padre/madre indicado (que podría ser nil)
func CrearPersona(nombre string, padre *Persona) *Persona {
	if padre != nil && padre.hij_mayor != nil && padre.hij_menor != nil {
		panic("En este modelo sólo permitimos hasta 2 hijos")
	}

	per := administrador.PedirMemoria[Persona]()
	(*per).nombre = nombre
	if padre == nil {
		return per
	}
	(*per).padre = padre

	if padre.hij_mayor == nil {
		padre.hij_mayor = per
	} else {
		padre.hij_menor = per
	}

	return per
}

// Imprimir imprime a todos los miembros de la familia
func (per *Persona) Imprimir() {
	if per == nil {
		return
	}
	fmt.Println(per.nombre)
	per.hij_mayor.Imprimir()
	per.hij_menor.Imprimir()
}

// Destruir libera la memoria (simulada) de esta Persona y todos sus descendientes
// 2T(n/2) + O(1)
// O(n log n)
func (per *Persona) Destruir() {
	fmt.Println(per)
	// Antes de la clase de AB
	// if per.hij_mayor == nil && per.hij_menor == nil {
	// 	administrador.LiberarMemoria[Persona](per)
	// 	return
	// }

	// if per.hij_mayor != nil {
	// 	per.hij_mayor.Destruir()
	// }
	// if per.hij_menor != nil {
	// 	per.hij_menor.Destruir()
	// }

	// Despues de clases de AB
	if per == nil {
		return
	}
	per.hij_mayor.Destruir()
	per.hij_menor.Destruir()
	
	administrador.LiberarMemoria[Persona](per)
}
