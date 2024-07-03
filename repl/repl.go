package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/confusedOrca/interpreter/lexer"
	"github.com/confusedOrca/interpreter/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lxr := lexer.New(line)

		for tkn := lxr.NextToken(); tkn.Type != token.EOF; tkn = lxr.NextToken() {
			fmt.Printf("%+v\n", tkn)
		}
	}
}
