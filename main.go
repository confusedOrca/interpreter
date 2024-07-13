package main

import (
	"fmt"
	"os"

	"github.com/confusedOrca/interpreter/repl"
)

var VERBOSE = true

func main() {
	fmt.Println("Type in commands:")
	repl.Start(os.Stdin, os.Stdout)
}
