package main

import (
	"fmt"
)

func main() {
	// Defininicion de un slice de enteros
	nums := []int{2, 1, 5, 4, 3}

	// mostrar el arreglo desordenado
	fmt.Println("Arreglo desordenado:")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}

	// ordenamiento del arreglo
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			// Si el elemento en la posicion i es menor que el elemento en la posicion j
			if nums[i] < nums[j] {
				var temp int = nums[j]
				// intercambiamos los elementos
				nums[j] = nums[i]
				nums[i] = temp
			}
		}
	}

	// mostrar el arreglo ordenado
	fmt.Println("\nArreglo ordenado:")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
}
