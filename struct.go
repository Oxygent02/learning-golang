package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) welcome(addMessage string) {
	fmt.Println("Hello,", customer.Name, "\n"+addMessage)
}

func main() {
	var pelanggan1 Customer
	pelanggan1.Name = "John Doe"
	pelanggan1.Address = "Jl Ikan Ikan"
	pelanggan1.Age = 24

	fmt.Println(pelanggan1)

	pelanggan2 := Customer{
		Name: "Jane Doe",
		Age:  22,
	}

	fmt.Println(pelanggan2)

	pelanggan3 := Customer{"Jimmy Doe", "", 0}
	fmt.Println(pelanggan3)

	pelanggan3.welcome("")
	pelanggan2.welcome("What do you need?")
}
