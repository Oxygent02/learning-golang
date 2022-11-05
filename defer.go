package main

import "fmt"

/*
* defer: seperti finally(), akan dijalankan diakhir program
* panic: seperti throw exception
* recover: mendapatkan message dari panic
 */

func main(){
	runApplication(true)
}

func catchError(){
	message := recover()
	if message != nil {
		fmt.Println("Error", message)
	}
}

func runApplication(isError bool){
	defer catchError()
	if isError {
		panic("Application Error")
	}

	fmt.Println("Application run well")
}