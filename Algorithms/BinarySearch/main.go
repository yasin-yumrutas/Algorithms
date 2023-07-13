package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 5, 7, 10, 24, 33, 45, 56, 61}
	target := 33

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Hedef değer %d, dizinin %d. indeksinde bulundu.\n", target, index)
	} else {
		fmt.Printf("Hedef değer %d dizide bulunamadı.\n", target)
	}
}
