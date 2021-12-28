package object

import (
	"strings"
)

const (
	ERROR_MULTIPLICATION_STRING = "La multiplication de <object.String> n'est pas supporté avec la valeur "
	ERROR_ADDITION_STRING       = "L'addition de <object.String> n'est pas supporté avec la valeur "
	ERROR_SUBTRACTION_STRING    = "La soustraction de <object.String> n'est pas supporté avec la valeur "
	ERROR_DIVISION_STRING       = "La division de <object.String> n'est pas supporté avec la valeur "
)

type String struct {
	Value string
}

// Type retourne le type de l'objet (OBJECT_STRING)
func (s String) Type() Type { return OBJECT_STRING }

// ToString retournes la représentation sous forme de string la valeur
// de l'obect (String)
func (s String) ToString() string { return s.Value }

// ToInterface retournes la valeur de l'objet (String)
func (s String) ToInterface() interface{} { return s }

// Add retournes la concaténation entre la valeur de l'objet présent et celui
// passé en paramètre
func (s String) Add(oth Object) Object {
	switch oth.(type) {
	case Number, String:
		return String{Value: s.Value + oth.ToString()}
	default:
		return Error{Error: ERROR_ADDITION_STRING + oth.ToString()}
	}
}

// Sub retournes une Erreur puisqu'elle n'est pas implémentée pour l'objet
// de type String
func (s String) Sub(oth Object) Object {
	return Error{Error: ERROR_SUBTRACTION_STRING + oth.ToString()}
}

// Mul retournes une répétition de n fois la chaîne de caractère si l'objet
// présent est de type (object.Number). Sinon retourne une erreur
func (s String) Mul(oth Object) Object {
	switch o := oth.(type) {
	case Number:
		return String{Value: strings.Repeat(s.Value, int(o.Value))}
	default:
		return Error{Error: ERROR_MULTIPLICATION_STRING + oth.ToString()}
	}
}

// Div retournes une Erreur puisque la division n'est pas supporté pour l'objet
// de type String
func (s String) Div(oth Object) Object {
	return Error{Error: ERROR_DIVISION_STRING + oth.ToString()}
}
