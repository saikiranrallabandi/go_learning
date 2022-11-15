package main
import (
	"fmt"
	"sync"
	"runtime"
)
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}


func main() {
	runtime.GOMAXPROCS(100)
	i := 0
	for i < 10 {
		i++
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello()  { 
	 fmt.Println("Hello #%v\n", counter)
	 m.RUnlock()
	 wg.Done()

}

func increment()  {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()

}
