package main

import (
	"fmt"
	"go/printer"
)

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
    w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}