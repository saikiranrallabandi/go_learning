// for dynamic allocation we use slice of slice

package main
import (
	"fmt"
)
func main() {
	var rows int
	var cols int
	rows = 7
	cols = 9
	var towdslices = make([][]int, rows)
	var i int
	for i = range towdslices {
		towdslices[i] = make([]int,cols)
	}
	fmt.Println(towdslices)

}