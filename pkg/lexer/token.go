package lexer

type Token uint8

const (
	ILLEGAL = iota
	EOF
	WS

	/* --- misc characters --- */

	SPACE
	AND

	/* --- keywords --- */

	WHOLE
	DECIMAL
)
