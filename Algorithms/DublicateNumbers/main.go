package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Println("Kaç sayı gireceksin?")
	fmt.Scan(&n)
	result := duplicateNumbersCalculation(n)

	for number, count := range result {
		fmt.Printf("Girilen sayı: %d // %d tane var\n", number, count)
	}
}

func duplicateNumbersCalculation(n int) map[int]int {
	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d. sayı = ", i+1)

		_, err := fmt.Scanln(&array[i])
		if err != nil {
			fmt.Println("Hatalı giriş tekrar deneyin")
			i--
			continue
		}

	}

	result := make(map[int]int)

	for _, num := range array {
		result[num]++
	}

	return result
}

//Niye fonksiyon içindeki hata bloğunu kullandığımı sorarsanız orda her program çalıştığında
//ilk atama default düşüyor onu düzeltmek için kullandım
