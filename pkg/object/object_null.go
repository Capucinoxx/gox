package object

type Null struct{}

func (n Null) Type() Type               { return OBJECT_NULL }
func (n Null) ToString() string         { return "null" }
func (n Null) ToInterface() interface{} { return n }
