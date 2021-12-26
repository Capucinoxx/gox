package object

import "bytes"

type Array struct{ Elements []Object }

// Type retourne le type de l'objet (OBJECT_ARRAY)
func (a Array) Type() Type { return OBJECT_ARRAY }

// ToString retournes la représentation sous forme de string la valeur
// de l'obect (Array)
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

// ToInterface retournes la valeur de l'objet (Array)
func (a Array) ToInterface() interface{} { return a }

// Add retournes un objet représentant la concaténation de l'array et celle
// passé en paramètre
func (a Array) Add(oth Object) Object {
	return Array{Elements: append(a.Elements, oth.(Array).Elements...)}
}

// Sub retournes une Erreur puisqu'elle n'est pas implémentée pour l'objet
// de type Array.
func (a Array) Sub(oth Object) Object {
	return Error{Error: "la soustraction n'est pas implémenté pour le type Array"}
}
