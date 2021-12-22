package object

import "fmt"

// Number objet représentant un nombre
type Number struct{ Value float64 }

func (i Number) Type() Type               { return OBJECT_NUMBER }
func (i Number) ToString() string         { return fmt.Sprint(i.Value) }
func (i Number) ToInterface() interface{} { return i }
