package main

import "fmt"

func main() {

	arr := []int{9, 2, 5, 7, 1, 4, 8, 6, 3}

	fmt.Println("Sıralanmamış dizi:", arr)
	mergeSort(arr)
	fmt.Println("Sıralanmış dizi:", arr)
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	// fmt.Printf("mid:%d\n", mid)
	left := arr[:mid]
	right := arr[mid:]

	mergeSort(left)
	mergeSort(right)

	merge(arr, left, right)
}

func merge(arr, left, right []int) {
	i := 0 // Sol parçanın indeksi
	j := 0 // Sağ parçanın indeksi
	k := 0 // Birleştirilmiş dizinin indeksi

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
