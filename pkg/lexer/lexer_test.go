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
		"space":             {c: `space`, tkn: &token.SPACE{}, literal: " "},
		"left parenthesis":  {c: `(`, tkn: &token.LPAREN{}, literal: "("},
		"right parenthesis": {c: `)`, tkn: &token.RPAREN{}, literal: ")"},
		"operand":           {c: `->`, tkn: &token.OPERAND{}, literal: "->"},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			scanner := lexer.NewScanner(strings.NewReader(tt.c))
			tkn, literal := scanner.Scan()

			if tt.tkn.GetToken() != tkn.GetToken() {
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
	input := `(file->keep "file.txt" ())`

	want := []struct {
		tkn     token.Token
		literal string
	}{
		{&token.LPAREN{}, "("},
		{&token.FILE{}, "file"},
		{&token.OPERAND{}, "->"},
		{&token.IDENTIFIER{}, "keep"},
		{&token.WS{}, " "},
		{&token.IDENTIFIER{}, "file.txt"},
		{&token.WS{}, " "},
		{&token.LPAREN{}, "("},
		{&token.RPAREN{}, ")"},
		{&token.RPAREN{}, ")"},
		{&token.EOF{}, ""},
	}

	scanner := lexer.NewScanner(strings.NewReader(input))

	for i, tt := range want {
		tkn, lit := scanner.Scan()

		if tt.tkn.GetToken() != tkn.GetToken() {
			t.Errorf("[%2d] -> token mismatch: exp=%q got=%q", i, tt.tkn, tkn)
		}

		if tt.literal != lit {
			t.Errorf("[%2d] -> literal mismatch: exp=%q got=%q", i, tt.literal, lit)
		}
	}
}
