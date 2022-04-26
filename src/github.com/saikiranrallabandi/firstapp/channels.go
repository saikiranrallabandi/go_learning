package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	    ch :=make(chan int, 50)
		wg.Add(2)// create two goroutines
		go func (ch <-chan int)  { // recive only channel
                for i := range ch {
					     fmt.Println(i)
				}


			    //  i := <- ch //data passing
				//  fmt.Println(i)
				//  i = <-ch
				//  fmt.Println(i)
				 wg.Done()

		}(ch)
		go func (ch chan<- int)  {
			      ch <- 27
			      ch <- 42 // arrow is pointing in to the channel
				  close(ch)
				  wg.Done()
		}(ch)
		wg.Wait()
}
