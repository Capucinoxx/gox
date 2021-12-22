package object

import "bytes"

type Array struct{ Elements []Object }

func (a Array) Type() Type { return OBJECT_ARRAY }
func (a Array) ToString() string {
	var out bytes.Buffer

	out.WriteRune('[')
	for i := 0; i < len(a.Elements)-1; i++ {
		out.WriteString(a.Elements[i].ToString())
		out.WriteString(", ")
	}
	out.WriteString(a.Elements[len(a.Elements)-1].ToString())
	out.WriteRune(']')

	return out.String()
}
func (a Array) ToInterface() interface{} { return a }
