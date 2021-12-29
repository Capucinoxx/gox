package object

type Chan struct {
	ch chan func() Object
	Fn func(obj Object) Object
}

func (c Chan) Type() Type               { return OBJECT_OPERATOR }
func (c Chan) ToString() string         { return "" }
func (c Chan) ToInterface() interface{} { return c }

// Make initalisation d'un cannal de fonction permettant l'ajout d'une fonction
// pour l'évaluation de l'expression
func (c *Chan) Make() {
	c.ch = make(chan func() Object, 1)
	c.ch <- func() Object { return Null{} }
}

func (c *Chan) GenFunc(str string) {
	switch str {
	case "+":
		c.Fn = func(obj Object) Object { return c.Add(obj) }
	case "-":
		c.Fn = func(obj Object) Object { return c.Sub(obj) }
	case "*":
		c.Fn = func(obj Object) Object { return c.Mul(obj) }
	}
}

// Eval fait l'évaluation de la composition de fonction présente dans le cannal
func (c *Chan) Eval() Object { return (<-c.ch)() }

// Add ajoute une fonction d'addition de l'objet (Object.Add) à la composition
// de fonction présente pour l'expression
func (c Chan) Add(obj Object) Object {
	old := <-c.ch
	c.ch <- func() Object { return old().Add(obj) }
	return nil
}

// Sub ajoute une fonction de soustraction de l'objet (Object.Sub) à la
// composition de fonction déjà présente pour l'expression
func (c Chan) Sub(obj Object) Object {
	old := <-c.ch
	c.ch <- func() Object { return old().Sub(obj) }
	return nil
}

func (c Chan) Mul(obj Object) Object {
	old := <-c.ch
	c.ch <- func() Object { return old().Mul(obj) }
	return nil
}

func (c Chan) Div(obj Object) Object {
	old := <-c.ch
	c.ch <- func() Object { return old().Div(obj) }
	return nil
}
