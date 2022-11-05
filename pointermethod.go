package main

import "fmt"

/*
Karena di Go Lang class nya merupakan pass by Value
maka ketika kita declare parameter dengan class yang sama
disarankan untuk menggunakan pointer, agar tidak banyak menggunakan memory
*/

type Man struct {
	Name string
}

func (man *Man) setName() {
	man.Name = "Mr. " + man.Name
}

func main() {
	boy := Man{"Boy"}
	boy.setName()
	fmt.Println(boy)
}
