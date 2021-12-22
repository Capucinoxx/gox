package lexer_test

import (
	"github.com/Capucinoxx/gox/pkg/lexer"
	"github.com/Capucinoxx/gox/pkg/token"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	cases := map[string]struct {
		c       string
		tkn     token.Token
		literal string
	}{
		"eof token":     {c: ``, tkn: &token.EOF{}},
		"illegal token": {c: `⿆`, tkn: &token.ILLEGAL{}, literal: `⿆`},
		"space char":    {c: ` `, tkn: &token.WS{}, literal: ` `},

		/* --- litteral --- */
		"identifier": {c: `toto`, tkn: &token.IDENTIFIER{}, literal: "toto"},

		/* --- misc characters --- */
		"left parenthesis":  {c: `(`, tkn: &token.LPAREN{}, literal: "("},
		"right parenthesis": {c: `)`, tkn: &token.RPAREN{}, literal: ")"},
		"operand":           {c: `->`, tkn: &token.OPERAND{}, literal: "->"},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			scanner := lexer.NewScanner(strings.NewReader(tt.c))
			tkn, literal := scanner.Scan()

			if tt.tkn.Token() != tkn.Token() {
				t.Errorf(
					"%q tkn mismatch: exp=%q got=%q <%q>",
					tt.c,
					tt.tkn,
					tkn,
					tt.literal,
				)
			}

			if tt.literal != literal {
				t.Errorf(
					"%q literal mismatch: exp=%q got=%q",
					tt.c,
					tt.literal,
					literal,
				)
			}
		})
	}
}

func TestInput(t *testing.T) {
	input := `(+ 1 2 3  6 "5")`

	want := []struct {
		tkn     token.Token
		literal string
	}{
		{&token.LPAREN{}, "("},
		{&token.OPERAND{}, "+"},
		{&token.WS{}, " "},
		{&token.IDENTIFIER{}, "1"},
		{&token.WS{}, " "},
		{&token.IDENTIFIER{}, "2"},
		{&token.WS{}, " "},
		{&token.IDENTIFIER{}, "3"},
		{&token.WS{}, "  "},
		{&token.IDENTIFIER{}, "6"},
		{&token.WS{}, " "},
		{&token.IDENTIFIER{}, "5"},
		{&token.RPAREN{}, ")"},
		{&token.EOF{}, ""},
	}

	scanner := lexer.NewScanner(strings.NewReader(input))

	for i, tt := range want {
		tkn, lit := scanner.Scan()

		if tt.tkn.Token() != tkn.Token() {
			t.Errorf("[%2d] -> token mismatch: exp=%d got=%d", i, tt.tkn.Token(), tkn.Token())
		}

		if tt.literal != lit {
			t.Errorf("[%2d] -> literal mismatch: exp=%q got=%q", i, tt.literal, lit)
		}
	}
}
