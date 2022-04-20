package main

import (
	"fmt"
)

func main() {
	    //  for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		// 	      fmt.Println(i,j)
		//  }
		for i := 0; i < 5; i++ {
		        fmt.Println(i)
				if i%2 == 0 {
					    i /=2
				} else {
					    i = 2*i + 1
				}


		}
}