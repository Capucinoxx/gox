package parser_test

import (
	"github.com/Capucinoxx/gox/pkg/parser"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	input := `(+ (+ 1 2 3) 1 2 3 4)`

	p := parser.NewParser(strings.NewReader(input))
	p.Parse()
}
