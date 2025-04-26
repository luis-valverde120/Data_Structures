# 📘 Busqueda Lineal

## 📖 Definición
El algoritmo de busqueda lineal o tambien conocido como busqueda secuencial es uno de los algoritmos mas simples para realizar busquedas. Este algoritmo recorre todo el arreglo hasta encontrar el valor dado.

---

## 🧠 Propiedades Clave
- Recorre todo el arreglo para encontrar un valor
- Si el valor a encontrar enta en el final del arreglo puede se tiene que recorrer todo el arreglo hasta el final para encontrar el valor
- Es facil de implementar
- No neceistan los datos estar ordenados

---

## 📦 Complejidad (Big O)
| Operación        | Mejor Caso | Promedio | Peor Caso |
|------------------|------------|----------|-----------|
| Búsqueda         | O(1)       | O(n)     | O(n)      |

---

## 🔍 ¿Cuándo usarlo?
Este algoritmo es util cuando se trabaja con datos que no estan ordenados, o datos en bruto los cuales no tienen un orden, para este tipo de casos el algiritmo de busqueda lineal es optimo.
- Datos pequeños y desordenados
- Datos almacenados en memoria contigua

---

## 📐 Analogías o Visualizaciones
Una analogia interesante es al buscar un libro en una biblioteca sin saber en que estanteria esta. Para nosotros encontrar este libro debemos revisar libro por libro hasta encontrarlo

---

## 🧵 Implementaciones

### 🔹 Python

```python
# Lineal search es una funcion que busca un elemento dentro de un arreglo
# y devuelve el indice del elemento si lo encuentra, o -1 si no lo enecuentra
def lineal_search(arr, target):
    # Recorrer el arreglo hasta encontrar el elemento objetivo
    for index, value in enumerate(arr):
        # si el elemento es igual al objetivo devolver el indice
        if value == target:
            return index

    # si no se encuentra el elemento devolver -1
    return -1
```

### 🔹 Go

``` go
func linealSearch(arr []int, findElement int) (int, bool) {
	// Recorrer el arreglo hasta encontrar el elemento
	for index, value := range arr {
		if value == findElement {
			return index, false // Retornar el índice del elemento encontrado
		}
	}
	return -1, true
}
```
