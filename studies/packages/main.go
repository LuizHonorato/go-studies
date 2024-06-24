package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	fmt.Println("Package main")

	error := checkmail.ValidateFormat("lhehonorato@gmail.com")
	fmt.Println(error)
}
