package main

import (
	"fmt"
)
// var email String = "Teemaporn@gmail.com"

func main() {
	// var name String = "Teemaporn"
	var age int = 21

	email := "Teemaporn@gmail.com"
	gpa := 3.46

	firstname, lastname := "Teemaporn", "Ruaengsri"

	fmt.Printf("name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age , email, gpa)

}