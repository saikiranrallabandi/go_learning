package main

import (
	"fmt"
)

func main()  {
	//Get User Input
	fmt.Printf("Please give me your name: ")
	var name string
	//name := ""
	fmt.Scanln(&name)
	fmt.Println("your name is", name)
}