package main

import (
	"fmt"
	"io"
)

func main() {
	     var w Writer = ConsoleWriter{}
		 w.Write([]byte("Hello Go!"))

		 myInt := IntCounter(0)
		 var inc Incrementor = &myInt // Incremntor to pointer object
		 for i := 0; i < 10; i++ {
			      fmt.Println(inc.Increment())

		 }

		 var myObj interface{} = NewBufferedWriterCloser()
		 if wc, ok := myObj.(WriterCloser); ok {
			    wc.Write([]byte("Hello youtube listeners, this is a test"))
				wc.Close()
		 }
		 r, ok := myObj.(io.Reader)
		 if ok {
			     fmt.Println(r)
		 } else {
			     fmt.Println("Conversation failed")
		 }

}

type Incrementor interface {
	    Increment() int
}

type IntCounter int

type Writer interface {
	Write([]byte) (int, error)
}
// type Writer interface {} empty interface
// }

type ConsoleWriter struct {}


func (cw ConsoleWriter) Write(data []byte) (int, error)  {
	   n, err := fmt.Println(string(data))
	   return n,err
}

func (ic *IntCounter) Increment() int {
	     *ic++
		 return int(*ic)
}