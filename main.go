package main

import "fmt"

type Person struct {
	name		string
	age			int
	profession	string
	isMale		bool
}

func main()  {
	var primitivo Person

	fmt.Printf("primitivo=%+v\n", primitivo)
}