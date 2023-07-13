package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	secim := rand.Intn(2)

	if secim == 0 {
		fmt.Println("Tura")
	} else {
		fmt.Println("Yazı")
	}
}
