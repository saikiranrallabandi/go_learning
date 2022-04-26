package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
    runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go _sayHello()
		m.Lock()
		go _increment()
	}

	// go _sayHello() // green Threads start
	// time.Sleep(100 * time.Millisecond)
	// msg := "Hello"
	// wg.Add(1)
	// go func (msg string)  {
	//         fmt.Println(msg)
	// 		wg.Done()
	// }(msg)
	//time.Sleep(100 * time.Millisecond) // nont best practises
	wg.Wait()
}

func _sayHello()  {
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func _increment()  {
	           counter++
			   m.Unlock()
			   wg.Done()
}