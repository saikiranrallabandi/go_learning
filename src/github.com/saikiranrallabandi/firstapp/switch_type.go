package main

import (
	"fmt"
)

func main() {
	     var i interface{} = true
		 switch i.(type) {
		case int:
			     fmt.Println("i is an integer")
		case string:
			    fmt.Println("I as a string")
		default:
		        fmt.Println("I don't know what I is")
		 }
}