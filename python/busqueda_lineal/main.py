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

if __name__ == "__main__":
    # Definicion de un arreglo y elemento a buscar
    arr = [1, 3, 2, 4, 6, 7, 5, 10, 9, 8]
    target = 5

    # llamada a la funcion lineal_search
    result = lineal_search(arr, target)

    if result == -1:
        print(f"El elemento {target} no se encuentra en el arreglo")
    else:
        print(f"El elemento {target} se encuentra en el indice {result}")