package object

type Type uint8

const (
	OBJECT_NULL Type = iota
	OBJECT_ERROR

	OBJECT_ARRAY

	OBJECT_NUMBER
	OBJECT_STRING
)

// Object représentations des différentes fonctions communes aux objets
// du langage.
type Object interface {
	// Type retourne le type de l'objet (Object)
	Type() Type

	// ToString retournes la représentation sous forme de string la valeur
	// de l'obect (Object)
	ToString() string

	// ToInterface retournes la valeur de l'objet (Object)
	ToInterface() interface{}

	// Add retournes une objet représentant l'addition de l'objet courant
	// et de l'objet en paramètre
	Add(Object) Object

	// Sub retournes une objet représentant la soustraction de l'objet courant
	// et de l'objet en paramètre
	Sub(Object) Object
}
