package main

import "fmt"

/*
*
INTERFACE
*/
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello 1", hasName.GetName())
}

type HasName2 interface {
	GetName() string
}

func SayHello2(hasName2 HasName2) {
	fmt.Println("Hello 2", hasName2.GetName())
}

/**
STRUCT
*/

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	macan := Animal{"macan"}
	SayHello(macan)
	SayHello2(macan)
}
