package main

import (
	"fmt"
	"learn-go-basic/database"
	// _ "learn-go-basic/database" // _ (blank identifier) untuk import package, tp package tidak terpakai
)

func main(){
	dbName := database.GetDatabase()
	fmt.Println(dbName)
}