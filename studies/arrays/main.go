package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays & Slices")

	var array1 [5]int
	array1[0] = 0
	array1[1] = 1
	array1[2] = 2
	array1[3] = 3
	array1[4] = 4
	fmt.Println(array1)

	array2 := [2]string{"P1", "P2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2}
	fmt.Println(slice)

	slice = append(slice, 3)
	fmt.Println(slice)
}
