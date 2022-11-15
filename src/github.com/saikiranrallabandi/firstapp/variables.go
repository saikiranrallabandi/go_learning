package main
import (
	"fmt"
)

// you can use like this k := 34


var (
	actorName    string = "saikiran"
	companyName  string = "Apple"
	doctorNumber int    = 3
	season       int    = 11
)


func main() {
	//var i int
	//i = 42 // not initialze
	var j float32 = 27
	//k := 45
	fmt.Printf("%v, %T", j, j)
}


// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 24
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
