package main

import (
	"fmt"
)

func main() {
	var T float64 = 20
	var r float64 = 4
	luasPermukaan := 2 * 3.14 * r * (r + T)
	fmt.Println("Luas permukaan tabung:", luasPermukaan)
}
