package main

import "fmt"

func main() {
	var nome string = "Luiz"
	sobrenome := "Honorato"

	fmt.Println(nome, sobrenome)

	var (
		address      string = "Rua Salvador Neves, 716"
		neighborhood string = "Vila Maria Luiza"
	)

	fmt.Println(address, neighborhood)

	city, state := "Ribeirão Preto", "SP"

	fmt.Println(city, state)

	const country string = "Brasil"
	fmt.Println(country)

	newCity := "Cravinhos"

	city = newCity

	fmt.Println(city)
}
