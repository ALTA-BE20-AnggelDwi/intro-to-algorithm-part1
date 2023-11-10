package main

import (
	"fmt"
)

func main() {
	var hargaAwal float64 = 370000
	var diskon float64 = 15
	hargaDiskon := (diskon / 100) * hargaAwal
	hargaAkhir := hargaAwal - hargaDiskon
	fmt.Println("Harga baju setelah diskon: Rp.", hargaAkhir)
}
