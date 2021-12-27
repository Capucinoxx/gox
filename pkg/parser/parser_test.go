package parser_test

import (
	"github.com/Capucinoxx/gox/pkg/parser"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"sum integer 1":      {`(+ 1 2 3 4)`, "10"},
		"sum integer 2":      {`(+ 1 2 3 4 "5")`, "15"},
		"sum string":         {`(+ "t" "t" "t")`, "ttt"},
		"difference integer": {`(- 10 1 2 3)`, "4"},
		"multiple sum":       {`(+ 1 2 3 4)(+ "t" "t" "t")`, "10ttt"},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			p := parser.NewParser(strings.NewReader(tt.input))
			p.Parse()
			res := p.Result()

			if tt.want != res {
				t.Errorf(
					"%q parser mismatch: exp=%q got=%q",
					tt.input,
					tt.want,
					res,
				)
			}
		})
	}
}
