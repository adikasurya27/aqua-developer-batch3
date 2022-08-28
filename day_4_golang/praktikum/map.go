package main

import "fmt"

func main() {
	var toko map[string]int 
	toko = map[string]int{} //inisialisasi

	toko["telur"] = 14000
	toko["beras"] = 12000
	toko["mie"] = 3000
	toko["bumbu"] = 1500
	toko["sayur"] = 2000
	toko["ayam"] = 35000

	fmt.Println(toko)
	fmt.Println("Harga ayam: ", toko["ayam"])
 }