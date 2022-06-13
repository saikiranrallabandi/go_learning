package main
import (
	"log"
	"os"
)
func main() {

	if len(os.Args) !=1 {
		log.Fatal("Fatal: Hellow world!")
	}
	log.Panic("Panic: Hello World")
}