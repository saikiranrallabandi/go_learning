package main
import (
	"fmt"
)
func main() {
	//Initialization of an array
	var arr = [5]int{1, 2, 4, 5, 6}
	var i int
	for i=0; i< len(arr); i++ {
		fmt.Println("priniting elements ", arr[i])
	}
	var value int
	for i, value = range arr{
		fmt.Println(" range ", value)
	}

	for _, value = range arr{
		fmt.Println(" blank range", value)
	}

}