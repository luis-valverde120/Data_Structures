## Ordenamiento de un arreglo python
En el ejercico realizado dobre el algoritmo para ordenar un arreglo de manera ascendente se utilzo una `list` con valores enteros desordenados en el cual de realizara el reordenamiento de menor a mayor realizando las comparaciones e intercambios para realizar el ordenamiento.

## Pasos del algoritmo

- **Primer ciclo (`i`)**: Recorre los elementos del arreglo
- **Segundo ciclo (`j`)**: Compara el elemento del ciclo `nums[i]` con el del ciclo `nums[j]` 
- **Comparacion de los elementos:** Condicion encardaga de comparar si `nums[i] < nums[j]` en caso de ser verdadera se intercambian los valores
- **Intercambio de valores**: 
  - *Decalracion de variable temporal* Esta variable llamada `temp` almacena el valor de `nums[j]`
  - *Intercambio de posiciones:* Para el intercambio de posiciones se utiliza el valor temporal y se realiza de la siguiente forma `nums[j] = nums[i]` y `num[i] = temp`

## Funcionamiento
En el funcionamiento de este algoritmo tenemos un arreglo el cual es
`[5, 2, 9, 1, 10, 6]`
Para ordenarlo se realizara 2 iteraciones en el arreglo y se realizara la verificacion de la posicion de cada elemento y finalmente el intercambio del los elementos

```python
# creacion del arreglo desordenado
    nums = [5, 2, 9, 1, 10, 6]
# ordenacion del arreglo
    for i in range(len(nums)):
        for j in range(0, len(nums)):
            if nums[i] < nums[j]:
                temp = nums[j]
                nums[j] = nums[i]
                nums[i] = temp
```