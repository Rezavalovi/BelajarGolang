package main

import "fmt"

func main() {
	const (
		firstName string = "Reza"
	 	lastName = "Valovi"

	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//error  -> tidak boleh diubah
	// firstname = "eza"
	// lastname = "lovi"
}