package parser

import (
	"fmt"
	"strings"

	globalstate "github.com/confusedOrca/interpreter/global_state"
)

var traceLvl int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLvl-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLvl += 1 }
func decIdent() { traceLvl -= 1 }

func trace(msg string) string {
	if globalstate.VERBOSE {
		incIdent()
		tracePrint("BEGIN " + msg)
		return msg
	} else {
		return ""
	}
}

func untrace(msg string) {
	if globalstate.VERBOSE {
		tracePrint("END " + msg)
		decIdent()
	}
}
