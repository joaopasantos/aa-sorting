package main

import (
	"fmt"

	"github.com/joaopasantos/aa-sorting/sorters"
)

func main() {
	// var array []int
	array := []int{6, 9, 4, 8, 7, 5, 10, 3, 2, 1}
	fmt.Println(array)
	// sorters.Insertion(array)
	sorters.Quick(array, 0, len(array)-1)
	fmt.Println(array)
}
