package main

import (
	"fmt"
)



func main(){
	days := [...]string{"Mon","Tue","Wed","Thu","Fri","Sat","Sun"}
	daySlice1 := days[3:4]
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice1[0] = "Danny"
	fmt.Println(daySlice1)

	daySlice1 = append(daySlice1, "1")
	for i := 0; i < len(daySlice1); i++{
		fmt.Println(daySlice1[i])
	}

	for index, element := range daySlice1 {
		fmt.Println(index, element)
	}


	sayHelloWithFilter("Anjing", spamFilter)


}


type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string{
	if(name == "Anjing"){
		return "..."
	}
	return name
}