package main

import "fmt"

func selectionSort(arrr []int) {
	n := len(arrr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arrr[j] < arrr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arrr[i], arrr[minIndex] = arrr[minIndex], arrr[i]
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Dizinin sırasız hali:", arr)

	selectionSort(arr)

	fmt.Println("Dizinin sıralı hali:", arr)
}

//NOT: dizi değişkenleri üzerinde kopyası verilen parametre üzerindeki değişiklik main değişkenide etkiler
//Önemli NOT: bu durum int vb. değişkenlerde geçerli değildir
//NOT: bu durum C dili içinde geçerlidir ama C# için geçerli değildir
