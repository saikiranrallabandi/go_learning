package main

import (
	"fmt"
)

func main() {
	    // for i := 0; i < 5; i++ {
		// 	      _sayMessage("Hello Sai", i)
		// }
		greeting := "Hello"
		name := "Stacey"
		_greeting(&greeting, &name)
		fmt.Println(name)
		s := _sum(1,2,3,4,5)
		fmt.Println("The Sum is", *s)
		d, err := _divide(5.0, 0.0)
		if err != nil {
			    fmt.Println(err)
				return
		}
		fmt.Println(d)
		func() {
			     fmt.Println("hello")
		}()
}

func _divide(a, b float64) (float64, error) {
	     if b == 0.0 {
			     return 0.0, fmt.Errorf("Cannot divide by Zero")
		 }
		 return a / b, nil
}

func _sum(values ...int) *int  {
	     fmt.Println(values)
		 result := 0
		 for _, v := range values {
			      result += v
		 }
		 return &result

}

func _sayMessage(msg string, idx int)  {
         fmt.Println(msg)
		 fmt.Println("The value of the index is", idx)

}

func _greeting(greeting, name *string){
	    fmt.Println(*greeting, *name)
		*name = "Ted"
		fmt.Println(*name)
}