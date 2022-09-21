package main
import (
	"fmt"
)
func main() {
	var lang = map[int]string {
		1: "Eng",
		2: "Fre",
	}
	var pro = make(map[int]string)
	pro[1] = "ch"
	pro[2] = "tab"

	var i int
	var value string
	for i, value = range lang {
		fmt.Println("lan", i, ":",value)
	}
	fmt.Println("product with key 2",pro[2])

}