package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	fmt.Println("Argumentos ->", args)
	columnas := validarInput(&args)

	var archivo *os.File
	if len(args) > 1 {
		ruta := args[1]
		archivo, err := os.Open(ruta)

		if err != nil {
			panic("Error: archivo fuente inaccesible")
		}

		defer archivo.Close()
	} else {
		archivo = os.Stdin
	}
	s := bufio.NewScanner(archivo)
	procesar(s, columnas)

}

func validarInput(args *[]string) int {
	if len(*args) > 2 {
		panic("Error: Cantidad erronea de parametros")
	}

	columnas, err := strconv.ParseInt((*args)[0], 0, 0)
	if err != nil {
		panic("Error: El numero de columnas no es un numero entero")
	}
	return int(columnas)
}

func procesar(s *bufio.Scanner, columnas int) {
	for s.Scan() {
		linea := s.Text()
		largo := len(linea)
		for i := 0; i <= largo/columnas; i++ {
			inicio := i * columnas
			var fin int
			if inicio+columnas > largo {
				fin = largo
			} else {
				fin = inicio + columnas
			}
			fmt.Println(linea[inicio:fin])
		}
	}

	err := s.Err()
	if err != nil {
		fmt.Println(err)
	}

}
