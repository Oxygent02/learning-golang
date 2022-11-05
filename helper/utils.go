package helper

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// function ini tidak bisa diakses dari luar
func sayHello(name string) {
	fmt.Println("Hello", name);
}