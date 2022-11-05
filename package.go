package main

import (
	"learn-go-basic/helper"
)

func main(){
	helper.SayHello("Boy")
	// helper.sayHello("Boy") // ini akan error, karena function dengan awal huruf kecil tidak bisa diakses dari luar
}