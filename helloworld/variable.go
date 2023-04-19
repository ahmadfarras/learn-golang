package main

import "fmt"

func main() {

	var name string = "Budianto"
	fmt.Println(name)

	var nameWithoutDataType = "Hansel"
	fmt.Println(nameWithoutDataType)

	nameWithoutVarAndDataType := "VarWithoutDataTypeAndVar"
	fmt.Println(nameWithoutVarAndDataType)

	var (
		firstname = "Firstname"
		lastname  = "Lastname"
	)

	fmt.Println(firstname, lastname)
}
