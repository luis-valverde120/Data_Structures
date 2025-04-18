## Ordenamiento de un arreglo desordenado
En el ejercico realizado sobre algoritmo tenemos que se realiza un ordenamiento ascendente de un arreglo que esta desordenado. Este toma como un `slice` de enteros el cual se realizaran los pasos necesarios para realizar el reordenamiento de menor a mayor realizando comparaciones e intercambios para realizar el ordenamiento.

## Pasos del algoritmo

- **Primer ciclo (`i`)**: Recorre los elementos del arreglo
- **Segundo ciclo (`j`)**: Compara el elemento del ciclo `nums[i]` con el del ciclo `nums[j]` 
- **Comparacion de los elementos:** Condicion encardaga de comparar si `nums[i] < nums[j]` en caso de ser verdadera se intercambian los valores
- **Intercambio de valores**: 
  - *Decalracion de variable temporal* Esta variable llamada `temp` almacena el valor de `nums[j]`
  - *Intercambio de posiciones:* Para el intercambio de posiciones se utiliza el valor temporal y se realiza de la siguiente forma `nums[j] = nums[i]` y `num[i] = temp`



## Funcionamiento
El funcionamiento del algoritmo es el siguiente
nosotros al tener un arreglo desordenado como el siguiente `[2, 1, 5, 4, 3]`
Para ordenarlo realizamos 2 iteraciones en el arreglo esto para y verificando cada posision del arreglo 


```go
// Defininicion de un slice de enteros
	nums := []int{2, 1, 5, 4, 3}

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
		}a
	}
```