package main

import (
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Prize int
}

type PruductSlice []Product

func (value PruductSlice) len() int {
	return len(value)
}

func (value PruductSlice) Less(i, j int) bool {
	return value[i].Prize < value[j].Prize
}

func (value PruductSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	products := []Product{
		{1, "Benih Lele", 50000},
		{2, "Pakan lele cap menara", 25000},
		{3, "Probiotik A", 75000},
		{4, "Probiotik Nila B", 10000},
		{5, "Pakan Nila", 20000},
		{6, "Benih Nila", 20000},
		{7, "Cupang", 5000},
		{8, "Benih Nila", 30000},
		{9, "Benih Cupang", 10000},
		{10, "Probiotik B", 10000},
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i] < products[j]
	})
	for _, value := range products {
		fmt.Println(value)
}
}