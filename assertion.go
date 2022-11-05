package main

import "fmt"

// assertion untuk casting tipe data dari tipe interface kosong
func main(){
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// bad example
	// var result2 interface{} = random()
	// var resultInt int = result2.(int) // akan kena panic karena ooops tidak bisa jadi int
	// fmt.Println(resultInt)

	// safe tricks
	switch result.(type) {
	case string:
		fmt.Println("string", result)
	case int:
		fmt.Println("int", result)
	default:
		fmt.Println("not decalred type")
	}

}

func random() interface{} {
	return false
}