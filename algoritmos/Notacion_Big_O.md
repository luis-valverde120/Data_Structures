## Notacion Big O
La **notación Big O**  o también llamada *notación de orden* es una herramienta utilizada en ciencias de la computación la cual es utilizada para describir **el comportamiento de un algoritmo conforme crece el tamaño de la entrada** `n`.

Esta notacion nos ayudara a entender como escalan los algortimo en términos de **tiempo de ejecucion** o incluso el **uso de memoria** esto tomando en cuenta la entrada `n` y como el algoritmo se comporta cuando crece la entrada. 

---

## ¿Por que se llama *Big O*?
Esta se denota por la letra **O** la cual proviene del **Orden de crecimiento** de un algortimo. ESta es una forma de **Expresar una cota superior  asintótica**, lo que describe un limite superior del peor comporamiento posible de una algiritmo en terminos de eficiencia.


---
### ejemplo
- `O(1)` → Tiempo constante, no importa el tamaño de la entrada.
- `O(n)` → Tiempo lineal, crece proporcionalmente a `n`.
- `O(log n)` → Tiempo logarítmico, muy eficiente para entradas grandes.
- `O(n^2)` → Tiempo cuadrático, crece rápido con entradas grandes.

---
## Tabla de complejidad de Big-O
| Notacion       | nombre          | Ejemplo  de uso                               |
|----------------|-----------------|-----------------------------------------------|
| `O(1)`         | Constante       | Acceder a un indice de un array               |
| `O(log(n))`    | Logaritmica     | Algoritmo de busqueda binaria                 |
| `O(log (n)^c)` | Polilogaritmica | Algoritmos que son eficiente es grafos        |
| `O(n)`         | Lineal          | Recorrido de un arreglo una sola vez          |
| `O(n log n)`   | Quasilineal     | Algortimos de ordenamiento eficiente          |
| `O(n^2)`       | Cuadratica      | Algoritmo de busqueda por pares en un arreglo |
| `O(n^c)`       | Polinomica      | Algunos algoritmos de fuerza bruta            |
| `O(C^n)`       | Exponencial     | Backtracking, problemas que son NP-complejos  |
| `O(n!)`        | Factorial       | Permutaciones exhautivas                      |

---
## Caracteristicas
- Esta describe el comportamiento de un algoritmo en funcion del crecimiento de su entrada `n`
- Muestra el **Limite superior** con respecto al tiempo de ejecucion o espacio requerido del algoritmo
- Siempre se enfoca en mostrar el **peor caso**
- Ignora constantes o terminos menores por ejemplo: `O(3n + 2)` se expresa solo como `O(n)`

---
## Importante
- Big o  **no mide el tiempo exacto**, sino mide la relacion que se tiene entre la entrada y los recursos usados
- Temos que tener en cuenta que no es lo mismo `O(n)` que `O(n log n)`.


