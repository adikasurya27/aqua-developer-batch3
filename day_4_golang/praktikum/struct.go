package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var dika Customer
	dika.Name = "Adika Surya Perdana"
	dika.Address = "Sidoarjo"
	dika.Age = 22

	fmt.Println(dika)

	var abi Customer
	abi.Name = "Arya Bintang Ilahi"
	abi.Address = "Sidoarjo"
	abi.Age = 21

	fmt.Println(abi)

	var rizal Customer
	rizal.Name = "Rizal Zamrun"
	rizal.Address = "Sidoarjo"
	rizal.Age = 21

	fmt.Println(rizal)
}