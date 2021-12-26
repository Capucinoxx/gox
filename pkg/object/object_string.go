package object

type String struct{ Value string }

func (s String) Type() Type               { return OBJECT_STRING }
func (s String) ToString() string         { return s.Value }
func (s String) ToInterface() interface{} { return s }
func (s String) Add(oth Object) Object {
	return String{Value: oth.(String).Value + s.Value}
}
