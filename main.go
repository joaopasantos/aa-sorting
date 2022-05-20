package main

import (
	"fmt"
	"os"

	"github.com/joaopasantos/aa-sorting/sorters"
)

func main() {
	var array []int
	array = inputNumbers(array)
	fmt.Println("- Vetor inserido:", array)

	showMenu()
	var method int
	var inputIsValid bool
	for !inputIsValid {
		inputIsValid = true
		fmt.Print("Insira o método de ordenação que deseja utilizar: ")
		fmt.Scanf("%d", &method)
		switch method {
		case 1:
			sorters.Bubble(array)
		case 2:
			sorters.Selection(array)
		case 3:
			sorters.Insertion(array)
		case 4:
			sorters.Quick(array, 0, len(array)-1)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Opção Inválida.")
			inputIsValid = false
		}
	}
	fmt.Println("- Vetor ordenado:", array)
}

func inputNumbers(array []int) []int {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Insira o número que ocupará a posição %d do vetor: ", i)
		var number int
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			panic(err)
		}
		array = append(array, number)
	}
	return array
}

func showMenu() {
	fmt.Println("\nMétodos de Ordenação")
	fmt.Println("[1] - Bubble Sort")
	fmt.Println("[2] - Selection Sort")
	fmt.Println("[3] - Insertion Sort")
	fmt.Println("[4] - Quick Sort")
	fmt.Println("[0] - Exit")
}
