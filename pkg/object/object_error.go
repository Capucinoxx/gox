package object

// Error object représentant une erreur
type Error struct{ Error string }

// Type retourne le type de l'objet (OBJECT_ERROR)
func (i Error) Type() Type { return OBJECT_ERROR }

// ToString retournes la représentation sous forme de string la valeur
// de l'obect (Error)
func (i Error) ToString() string { return i.Error }

// ToInterface retournes la valeur de l'objet (Error)
func (i Error) ToInterface() interface{} { return i }

// Add retournes l'erreur courante puisque l'addition d'erreur
// n'est pas implémenté
func (i Error) Add(oth Object) Object { return i }

// Sub retournes l'erreur courante puisque l'addition d'erreur
// n'est pas implémenté
func (i Error) Sub(oth Object) Object { return i }
