package token

import (
	"github.com/Capucinoxx/gox/pkg/object"
)

type LPAREN struct{ Node }

func (s *LPAREN) Token() uint8            { return TOKEN_LPAREN }
func (s *LPAREN) ToObject() object.Object { return object.Null{} }

type RPAREN struct{ Node }

func (s *RPAREN) Token() uint8            { return TOKEN_RPAREN }
func (s *RPAREN) ToObject() object.Object { return object.Null{} }

type OPERAND struct {
	Value string
	Node
}

func (s *OPERAND) Token() uint8   { return TOKEN_OPERAND }
func (s *OPERAND) String() string { return s.Value }
func (s *OPERAND) ToObject() object.Object {
	ch := object.Chan{}
	ch.GenFunc(s.Value)
	ch.Make()
	return ch
}
