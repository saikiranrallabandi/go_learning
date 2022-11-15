package main

import (
	"fmt"
)

func main()  {
	     statePopulations := make(map[string]int)
	     statePopulations = map[string]int{
			 "California" : 39250017,
			 "Texas": 27862596,
			 "Florida" : 20612439,
		 }
		 statePopulations["Georgia"] = 10232323
		 delete(statePopulations, "Georgia" )
		 //pop, ok := statePopulations["Texas"]
		 _, ok := statePopulations["California"]
		 fmt.Println(ok)
		 fmt.Println(len(statePopulations))



	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 	"California": 39250017,
	// 	"Texas":      27862596,
	// 	"Florida":    20612439,
	// }
	// for k, _ := range statePopulations {
	// 	fmt.Println(k)
	// }



}
