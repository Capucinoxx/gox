package token

import (
	"fmt"
	"github.com/Capucinoxx/gox/pkg/object"
)

const (
	TOKEN_ILLEGAL (uint8) = iota
	TOKEN_EOF
	TOKEN_WS

	/* --- litteral --- */

	TOKEN_IDENTIFIER

	/* --- misc characters --- */

	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_OPERAND
)

type Node struct {
	Parent *Node
	V      Token
	Value  string
	args   []*Node
}
type Token interface {
	Token() uint8
	ToObject() object.Object
}

type ILLEGAL struct{ Node }

func (s *ILLEGAL) Token() uint8 { return TOKEN_ILLEGAL }
func (s *ILLEGAL) ToObject() object.Object {
	return object.Error{
		Error: fmt.Sprintf("illegal token <%s>", s.Value),
	}
}

type EOF struct{ Node }

func (s *EOF) Token() uint8            { return TOKEN_EOF }
func (s *EOF) ToObject() object.Object { return object.Null{} }

type WS struct{ Node }

func (s *WS) Token() uint8            { return TOKEN_WS }
func (s *WS) ToObject() object.Object { return object.Null{} }
