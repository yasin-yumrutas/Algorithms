package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(100) + 1
	fmt.Println("1 ile 100 arasında bir sayı seçildi. Tahmin et!")

	for attempts := 1; ; attempts++ {
		var guess int
		fmt.Print("Tahmininiz: ")
		fmt.Scanln(&guess)

		if guess < target {
			fmt.Println("Daha yüksek bir sayı tahmin edin.")
		} else if guess > target {
			fmt.Println("Daha düşük bir sayı tahmin edin.")
		} else {
			fmt.Printf("Tebrikler! Doğru sayıyı %d tahminde buldunuz.\n", attempts)
			break
		}
	}
}
