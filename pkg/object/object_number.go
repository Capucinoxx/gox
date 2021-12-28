package object

import (
	"fmt"
	"strings"
)

const (
	ERROR_ADDITION_NUMBER    = "L'addition de <object.Number> n'est pas supporté avec la valeur "
	ERROR_SUBTRACTION_NUMBER = "L'addition de <object.String> n'est pas supporté avec la valeur "
)

// Number objet représentant un nombre
type Number struct{ Value float64 }

// Type retourne le type de l'objet (OBJECT_NUMBER)
func (i Number) Type() Type { return OBJECT_NUMBER }

// ToString retournes la représentation sous forme de string la valeur
// de l'obect (Number)
func (i Number) ToString() string { return fmt.Sprint(i.Value) }

// ToInterface retournes la valeur de l'objet (Number)
func (i Number) ToInterface() interface{} { return i }

// Add retournes une objet représentant l'addition du Nombre courant et celui
// passé en paramètre
func (i Number) Add(oth Object) Object {
	switch o := oth.(type) {
	case Number:
		return Number{Value: i.Value + o.Value}
	default:
		return Error{Error: ERROR_ADDITION_NUMBER + oth.ToString()}
	}
}

// Sub retournes une objet représentant la soustraction du Nombre courant et
// celui passé en paramètre
func (i Number) Sub(oth Object) Object {
	switch o := oth.(type) {
	case Number:
		return Number{Value: i.Value - o.Value}
	default:
		return Error{Error: ERROR_ADDITION_NUMBER + oth.ToString()}
	}
}

func (i Number) Mul(oth Object) Object {
	switch o := oth.(type) {
	case Number:
		return Number{Value: i.Value * o.Value}
	case String:
		return String{Value: strings.Repeat(o.Value, int(i.Value))}
	default:
		return Error{Error: ERROR_SUBTRACTION_NUMBER + oth.ToString()}
	}
}
