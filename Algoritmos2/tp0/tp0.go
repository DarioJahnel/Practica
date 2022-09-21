package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	aux := *y
	*y = *x
	*x = aux

}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	size := len(vector)
	if size == 0 {
		return -1
	}
	max := 0
	for i := 1; i < size; i++ {
		if vector[i] > vector[max] {
			max = i
		}
	}
	return max
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	size1 := len(vector1)
	size2 := len(vector2)
	max := size2
	if size1 > size2 {
		max = size1
	}

	for i := 0; i < max; i++ {
		if size1 <= i {
			return -1
		}
		if size2 <= i {
			return 1
		}
		if vector1[i] < vector2[i] {
			return -1
		}
		if vector1[i] > vector2[i] {
			return 1
		}
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := 0; i < len(vector); i++ {
		min := vector[i]
		pos := i
		for j := i + 1; j < len(vector); j++ {
			if min > vector[j] {
				min = vector[j]
				pos = j
			}
		}
		vector[pos] = vector[i]
		vector[i] = min
	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}

	return sumaRecursiva(&vector, len(vector)-1)
}

func sumaRecursiva(vector *[]int, size int) int {
	if size == 0 {
		return (*vector)[0]
	}

	return (*vector)[size] + sumaRecursiva(vector, size-1)
}

// EsPalindromo devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA.
func EsPalindromoIterativo(cadena string) bool {
	size := len(cadena)
	if size == 0 {
		return true
	}
	for i := 0; i < size/2; i++ {
		if cadena[i] != cadena[size-1-i] {
			return false
		}
	}
	return true
}

func EsPalindromo(cadena string) bool {
	if len(cadena) < 2 {
		return true
	}

	res := palRecursivo(&cadena, 0, len(cadena)-1)
	return res
}

func palRecursivo(cadena *string, str int, end int) bool {
	if (*cadena)[str] != (*cadena)[end] {
		return false
	}

	if str < end {
		return palRecursivo(cadena, str+1, end-1)
	}

	return true
}
