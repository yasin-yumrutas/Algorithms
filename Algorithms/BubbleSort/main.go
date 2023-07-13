package main

import (
	. "fmt"
)

func bubbleSort(arr *[]int) []int {
	for i, v := range *arr {
		Printf("%d.index==%d.\n", i, v)
	}

	Printf("\n\n\n")

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr); j++ {
			if (*arr)[i] < (*arr)[j] {
				temp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}

	for i, v := range *arr {
		Printf("%d.index==%d.\n", i, v)
	}
	Printf("\n\n\n")
	return *arr
}

func main() {
	// bubble := []int{8, 9, 5, 7, 6, 4, 1, 2, 3}
	var bubble = []int{8, 9, 5, 7, 6, 4, 1, 2, 3}
	arr := bubbleSort(&bubble)

	for i, v := range arr {
		Printf("%d.index==%d.\n", i, v)
	}

	Printf("\n\n\n")
	for i, v := range bubble {
		Printf("%d.index==%d.\n", i, v)
	}
}

//Chatgpt çözümü

/*
package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap arr[j] and arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    // Örnek bir dizi oluşturalım
    arr := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println("Dizinin sırasız hali:", arr)

    bubbleSort(arr)

    fmt.Println("Dizinin sıralı hali:", arr)
}

*/
