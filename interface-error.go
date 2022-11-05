package main

import (
	"errors"
	"fmt"
)

func divide(value int, divider int) (int, error){
	if divider == 0{
		return 0, errors.New("not valid divider")
	}
	return value/divider, nil
}

func main(){
	var errorExample error = errors.New("Wow Error")
	fmt.Println(errorExample)

	result, msgError := divide(4, 0)
	fmt.Println(result, msgError)
}