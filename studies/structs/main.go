package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var user1 user
	user1.name = "Luiz"
	user1.age = 32

	fmt.Println(user1)

	user2 := user{"Fernanda", 31}
	fmt.Println(user2)

	user3 := user{name: "Elisa"}
	fmt.Println(user3)
}
