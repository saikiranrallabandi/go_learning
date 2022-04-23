package main

import (
	"fmt"
	"log"
	"net/http"
)

func http_(){
        http.HandleFunc("/", func(w  http.ResponseWriter, r *http.Request){
		        w.Write([]byte("Hello go!"))
		})
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
		        panic(err.Error())
		}
}





func main() {
	//http_()
	fmt.Println("start")
	defer func() {
		     if err := recover(); err != nil {
				    log.Println("Error:", err)
			 }
	}()
	panic("something bad happened")
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
}