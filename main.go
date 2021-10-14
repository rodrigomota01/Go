package main

import "fmt"

type Person struct {
	name		string
	age			int
	profession	string
	isMale		bool
}

func main()  {
	primitivo := Person {
		name: "Rodrigo",
		age: 19,
		profession: "Analista",
		isMale: true,
	}

	fmt.Printf("primitivo=%+v\n", primitivo)
}