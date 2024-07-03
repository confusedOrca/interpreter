package parser

import (
	"testing"

	"github.com/confusedOrca/interpreter/parser"
)

func checkParserErrors(t *testing.T, parser *parser.Parser) {
	errors := parser.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
