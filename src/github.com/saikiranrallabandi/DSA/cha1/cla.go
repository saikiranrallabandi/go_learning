package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("Need one or more argument")
		return
	}

	var min, max float64
	for i :=1; i < len(argument); i++ {
		n, err := strconv.ParseFloat(argument[i], 64)
		if err!= nil {
			continue
		}
		if i ==1 {
			min = n
			max = n
			continue
		}

	if n < min {
            min = n
        }
        if n > max {
            max = n
        }

    fmt.Println("Min:", min)
    fmt.Println("Max:", max)
	}
}