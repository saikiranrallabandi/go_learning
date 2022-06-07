package main

import (
	"fmt"
	"log"
	"os"
	"path"
)
func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
	// The call to os.OpenFile()  creates the log file for writing,

	// if it does not already exists, or opens it for writing

	// by appending new data at the end of it ( os.0_Appending)
	if err !=nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Hello there!")
	iLog.Panicln("Mastering Go 3rd edition!")
}