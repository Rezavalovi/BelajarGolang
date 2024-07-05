package main

import "fmt"

func main() {

	name := "Reza Valovi"  //deklarasi awal
	fmt.Println(name)

	name = "Valovi Reza"
	fmt.Println(name)

	var (
		firstName = "Reza"
		lastName = "Valovi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}