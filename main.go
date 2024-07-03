package main

import (
	"fmt"
	"os"

	"github.com/confusedOrca/interpreter/repl"
)

func main() {
	fmt.Println("Type in commands:")
	repl.Start(os.Stdin, os.Stdout)
}
