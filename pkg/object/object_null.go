package object

type Null struct{}

// Type retourne le type de l'objet (OBJECT_NULL)
func (n Null) Type() Type { return OBJECT_NULL }

// ToString retournes la représentation sous forme de string la valeur
// de l'obect (Null)
func (n Null) ToString() string { return "null" }

// ToInterface retournes la valeur de l'objet (Null)
func (n Null) ToInterface() interface{} { return n }

// Add retournes l'objet passé en paramètre
func (n Null) Add(oth Object) Object { return oth }

// Sub retournes l'objet passé en paramètre
// TODO: doit envoyé  0 - valeur si nombre
func (n Null) Sub(oth Object) Object { return oth }

func (n Null) Mul(oth Object) Object { return oth }

func (n Null) Div(oth Object) Object { return oth }
