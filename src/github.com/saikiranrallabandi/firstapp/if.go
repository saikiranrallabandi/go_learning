package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
			 "California" : 39250017,
			 "Texas": 27862596,
			 "Florida" : 20612439,
	}
	if pop, ok := 	statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}
}