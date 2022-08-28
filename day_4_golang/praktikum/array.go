package main

import "fmt"

func main() {

	var city [5]string

	city[0] = "Sidoarjo"
	city[1] = "Jember"
	city[2] = "Surabaya"
	city[3] = "Malang"
	city[4] = "Pasuruan"

	fmt.Println(city[0])
	fmt.Println(city[1])
	fmt.Println(city[2])
	fmt.Println(city[3])
	fmt.Println(city[4])

	fmt.Println(len(city))
}