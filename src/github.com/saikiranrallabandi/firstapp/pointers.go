package main

import (
	"fmt"
)

type myStruct struct {
	    foo int
}


func _struc ()  {
	       var ms *myStruct
		   ms = new(myStruct)
		   (*ms).foo = 42
		   fmt.Println((*ms).foo)
}

func _varIn()  {
	       a := [3]int{1,2,3}
		   b := a
		   fmt.Print(a,b)
		   a[1] = 42
		   fmt.Print(a,b)

}

func _map()  {
	     a := map[string]string {"foo": "bar", "baz" : "buz"}
		 b := a
		 fmt.Print(a,b)
		 a["foo"] =  "qux"
		 fmt.Print(a,b)

}



func main() {
	    _map()
	    _varIn()
	    _struc()
        var a int =  42
		var b *int = &a
		fmt.Println(a, *b)
}