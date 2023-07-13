package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	data := make([]byte, 1000)
	for i := 0; i < len(data); i++ {
		data[i] = byte(rand.Intn(256))
	}

	copyData := make([]byte, 300-200+1)
	copy(copyData, data[200:301])
	for i := 0; i < 400; i++ {
		data[i+400] = copyData[rand.Intn(len(copyData))]
	}

	fmt.Println("200-300 arasındaki veriler:")
	for i := 200; i <= 300; i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()
	fmt.Println("400-700 arasındaki verilere dağıtılan veriler:")
	for i := 400; i <= 700; i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()
}
