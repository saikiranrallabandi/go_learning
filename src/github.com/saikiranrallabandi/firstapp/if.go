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

	    number :=50
      	guess := 30
		if guess < 1 || guess > 100 {
			fmt.Println("The guess must be between 1 and 100!")
		}
		if guess >=1 && guess <= 100 {
			if guess < number {
				fmt.Println("Too Low")
			}
			if guess > number {
				fmt.Println("Too high")
			}
			if guess == number {
                 fmt.Println("you got it")
			}
			fmt.Println(number >= guess, number != guess)
		}
}