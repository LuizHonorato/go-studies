package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	fmt.Println(somar(2, 7))

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Func F 2")

	resultSum, resultSub := calculos(10, 20)
	fmt.Println(resultSum, resultSub)
}
