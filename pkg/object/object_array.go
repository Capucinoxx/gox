package object

import (
	"bytes"
)

const (
	ERROR_SUBTRACTION_ARRAY = "La soustraction de <object.Array> n'est pas supporté avec la valeur "
)

type Array struct{ Elements []Object }

// Type retourne le type de l'objet (OBJECT_ARRAY)
func (a Array) Type() Type { return OBJECT_ARRAY }

// ToString retournes la représentation sous forme de string la valeur
// de l'obect (Array)
func (a Array) ToString() string {
	var out bytes.Buffer

	out.WriteRune('[')
	if len(a.Elements) > 0 {
		for i := 0; i < len(a.Elements)-1; i++ {
			out.WriteString(a.Elements[i].ToString())
			out.WriteString(", ")
		}
		out.WriteString(a.Elements[len(a.Elements)-1].ToString())
	}

	out.WriteRune(']')

	return out.String()
}

// ToInterface retournes la valeur de l'objet (Array)
func (a Array) ToInterface() interface{} { return a }

// Add retournes un objet représentant la concaténation de l'array et celle
// passé en paramètre
func (a Array) Add(oth Object) Object {
	switch o := oth.(type) {
	case Array:
		return Array{Elements: append(a.Elements, o.Elements...)}
	case Null:
		return a
	default:
		return Array{Elements: append(a.Elements, oth)}
	}

}

// Sub retournes une Erreur puisqu'elle n'est pas implémentée pour l'objet
// de type Array.
func (a Array) Sub(oth Object) Object {
	switch oth.(type) {
	case Null:
		return a
	default:
		return Error{Error: ERROR_SUBTRACTION_ARRAY + oth.ToString()}
	}
}
