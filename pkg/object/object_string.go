package object

type String struct{ Value string }

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
	return String{Value: oth.(String).Value + s.Value}
}

// Sub retournes une Erreur puisqu'elle n'est pas implémentée pour l'objet
// de type String
func (s String) Sub(oth Object) Object {
	return Error{Error: "la soustraction n'est pas implémenté pour le type String"}
}
