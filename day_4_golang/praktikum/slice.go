package main

import "fmt"

func main() {
	var days = [7]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	var slice1 = days[1:4]
	fmt.Println(slice1)
}