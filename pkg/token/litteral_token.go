package token

import (
	"github.com/Capucinoxx/gox/pkg/object"
	"strconv"
)

type IDENTIFIER struct {
	Value string
	Node
}

func (s *IDENTIFIER) Token() uint8 { return TOKEN_IDENTIFIER }
func (s *IDENTIFIER) ToObject() object.Object {
	v, err := strconv.ParseFloat(s.Value, 64)
	if err != nil {
		return object.String{
			Value: s.Value,
		}
	}
	return object.Number{
		Value: v,
	}
}
