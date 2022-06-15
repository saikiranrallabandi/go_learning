package main
import (
	"fmt"
)

// “Let's say you have an IProcessor interface with a process method,
//the Adapter class implements the process method and has an Adaptee instance as an attribute. T
//The Adaptee class has a convert method and an adapterType instance variable.
//The developer while using the API client calls the process interface method to invoke convert on Adaptee


//IProcess interface
type IProcess interface {
	process()
}

//Adapter struct

type Adapter struct {
	adaptee Adaptee
}

//Adpater class method process
func(adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

//Adaptee Struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert()  {
	fmt.Println("Adaptee convert method")

}


func main() {
	var process IProcess = Adapter{}
	process.process()
}