package object

type Type uint8

const (
	OBJECT_NULL (uint8) = iota
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
}
