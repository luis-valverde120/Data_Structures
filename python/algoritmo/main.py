
# Algoritmo de ordenacion 
def ordenar(nums):
    # ordenacion del arreglo
    for i in range(len(nums)):
        for j in range(0, len(nums)):
            if nums[i] < nums[j]:
                temp = nums[j]
                nums[j] = nums[i]
                nums[i] = temp

    return nums
if __name__ == "__main__":
    # creacion del arreglo desordenado
    nums = [5, 2, 9, 1, 10, 6]
    print("Arreglo desordenado:", nums)
    nums = ordenar(nums)
    print("Arreglo ordenado:", nums)
