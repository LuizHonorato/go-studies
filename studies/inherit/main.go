package main

import "fmt"

type pearson struct {
	firstname string
	lastname  string
	age       uint8
	height    uint8
}

type student struct {
	pearson
	course     string
	university string
}

func main() {
	fmt.Println("Herança")

	p1 := pearson{"Luiz", "Honorato", 32, 184}
	fmt.Println(p1)

	e1 := student{p1, "Golang", "Udemy"}
	fmt.Println(e1.firstname)
}
