package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 5, 7, 10, 24, 33, 45, 56, 61}
	target := 33

	index := linearSearch(arr, target)

	if index != -1 {
		fmt.Printf("Hedef değer %d, dizinin %d. indeksinde bulundu.\n", target, index)
	} else {
		fmt.Printf("Hedef değer %d dizide bulunamadı.\n", target)
	}
}
