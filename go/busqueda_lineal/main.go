package main

import (
	"fmt"
)

func linealSearch(arr []int, findElement int) int {
	// Recorrer el arreglo hasta encontrar el elemento
	for index, value := range arr {
		if value == findElement {
			return index // Retornar el índice del elemento encontrado
		}
	}
	return -1
}

func main() {
	// Definir el arreglo y el elemento a buscar

	arr := []int{5, 3, 8, 6, 2, 7}

	findElement := 6

	// Llamar a la funcion de busqueda lineal
	value := linealSearch(arr, findElement)

	if value != -1 {
		fmt.Printf("Elemento %d encontrado en el índice %d\n", findElement, value)
	} else {
		fmt.Printf("Elemento %d no encontrado\n", findElement)
	}
}
