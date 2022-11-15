// package main
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )
// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	robots, err := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello sai"))

	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	"fmt"
	"log"
)

-----------------------------------
Anonamus function
-----------------------------------

func main() {
	fmt.Println("start")
	defer func() { // Anonmus func
		if err := recover()// return nil if application is panic; err != nil {
			log.Println("Error:", err)
		}
	}() // execute
	panic("Something happened")
	fmt.Println("end")
}


-----------------------------------
Pointers
-----------------------------------

func main() {
	a := [3]int{1,2,3}
	b := &a[0]
	c := &a[1] - 4
	fmt.Println("%v %p %p\n", a, b, c)
}


----------------------------------------
Interfaces implicit
----------------------------------------
package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
