package diccionario_test

import (
	TDADiccionario "abb/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearDiccionarioOrdenado[int, int]()
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
}

func TestDiccionarioUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearDiccionarioOrdenado[int, int]()
	dic.Guardar(10, 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(10))
	require.Equal(t, 10, dic.Obtener(10))
	// require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(20) })
	// require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(20) })
	require.Equal(t, 10, dic.Borrar(10))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
}
