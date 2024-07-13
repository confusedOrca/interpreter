package repl

import (
	"bufio"
	"fmt"
	"io"

	globalstate "github.com/confusedOrca/interpreter/global_state"
	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/parser"
)

const PROMPT = ">>"

var VERBOSE = true

func Start(in io.Reader, out io.Writer) {
	globalstate.VERBOSE = false
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lxr := lexer.New(line)
		parser := parser.New(lxr)
		program := parser.ParseProgram()

		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parsing Errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
