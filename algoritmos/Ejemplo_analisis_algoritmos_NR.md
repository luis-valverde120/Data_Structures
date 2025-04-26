# 📘 Ejemplos analisis algoritmico (Iterativo / No recursivo)

## 📖 Definición
El [analisis de algoritmos](./Analisis_de_algoritmos.md) permite encontrar una sintota que muestra el comportamiento del algoritmo. 

---

## 🧪 Ejercicios Recomendados
Ejemplos de como se pueden identificar la complejidad algoritmica no recursiva

- **_Complejidad constante_**

Esta se denota en big O como n(1) como ejemplo de esto tenemos.

```python
x = 1 # O(1)
y = x = 2 # O(1)
print("hola mundo") # O(1)
lista = [1, 2, 3, 4] # O(1)
```
> No depende de la entrada

- **_Complejidad Logaritmica_**

La complejidad logaritmica se da cuando se descarta en cada iteracion la mitad de los elementos. **No se accede a todos los elementos del arreglo**

```go
func binary_search(arr []int, target) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        // asignar la mitad del valor
        mid := math.Round((low + high) / 2) // O(n/2)  
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1 // Reasignar el valor de low para cuando el targer es mayor
        } else {
            high = mid + 1 // Reasignar el valor de high en caso de que target sea menor 
        }
    }
}

```

Se puede ver como al buscar los datos se va reduciendo a la mitad la entrada en cada iteracion, esto nos indica que nuestra funcion va estar definida por `log`, otro factor a tomar en cuenta es que no se accede a todos los elementos del arreglo el algoritmo se puede ver de la siguiente manera

```math
n/2 \rightarrow n/4 \rightarrow \cdots = log(n) 
```

- **_Complejidad lineal_**

La complejidad lineal se puede encontrar cuando en un arreglo accedemos a cada elemento dentro de un arreglo. Si suponemos que tenemos `n` cantidad de elementos y si vamos a recorrer cada elemento esta complejidad depende de la cantidad de elementos de la lista `n` 

```python
lista = [1, 2, 3, 4, 5, 6] # O(1)

for elemento in lista: # O(n)
    print(elemento) 
```
>En big O siempre se toma el orden peor

```math
O(1) + O(n) = O(n)
```

- **_Complejidad N^2_**
Esta complejidad se puede entender por el recorrido del arreglo 2 veces. aqui se realiza la iteracion del arreglo una sola vez nos da `O(n)`, pero dentro de cada iteracion nosotros volvemos a realizar otra iteracion `O(n)` en este caso tenemos

```math
O(n)*O(n) = O(n^2)
```

```go
for i := range arr { // O(n)
    for j := range arr { // O(n)
        print(i + j)
    }
}
```
