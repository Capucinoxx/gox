package token

type Token uint8

const (
	ILLEGAL = iota
	EOF
	WS

	/* --- litteral --- */

	IDENTIFIER

	/* --- misc characters --- */

	SPACE
	AND
	LPAREN
	RPAREN
	OPERAND

	/* --- keywords --- */

	FILE
	WHOLE
	DECIMAL
)
