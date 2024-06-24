package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variable1 int = 10

	var variable2 int = variable1

	fmt.Println(variable1, variable2)
	variable1++
	fmt.Println(variable1, variable2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variable3 int
	var pointer *int

	fmt.Println(variable3, pointer)

	variable3 = 100
	pointer = &variable3
	fmt.Println(variable3, *pointer)
}
