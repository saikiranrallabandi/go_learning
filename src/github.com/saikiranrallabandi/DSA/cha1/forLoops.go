package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i*i, " ")
	}
	fmt.Println()

	i := 0
	for ok := true; ok; ok = (i !=10) {
		fmt.Println(i*i, " ")
		i++

	}
	fmt.Println()

	// This is a slice but range also works with arrays

	arrays := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range arrays {
		fmt.Println("index:", i, "value:", v)
	}
}