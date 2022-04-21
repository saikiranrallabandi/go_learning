package main

import (
	"fmt"
)


func main() {
//Loop:
	    //  for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		// 	      fmt.Println(i,j)
		//  }
		// for i := 0; i < 5; i++ {
		//         fmt.Println(i)
		// 		if i%2 == 0 {
		// 			    i /=2
		// 		} else {
		// 			    i = 2*i + 1
		// 		}


		// }
		// i :=0
		// for ;i < 5; {
		// 	      fmt.Println(i)
		// 		  i++
		// }
		// for i := 1; i <=3; i++ {
		// 	for j := 1; j <=3; j++ {
		// 		    fmt.Println( i * j)
		// 			if i * j >= 3 {
		// 				   break Loop
		// 			}

		// 	}

		// }
		s := []int{1,2,3}
		fmt.Println(s)
		for k, v := range s {
			     fmt.Println(k,v)
		}
}