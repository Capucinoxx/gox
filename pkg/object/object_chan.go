package object

type Chan struct{ ch chan func() Object }

// Make initalisation d'un cannal de fonction permettant l'ajout d'une fonction
// pour l'évaluation de l'expression
func (c *Chan) Make() {
	c.ch = make(chan func() Object, 1)
	c.ch <- func() Object { return Null{} }
}

// Eval fait l'évaluation de la composition de fonction présente dans le cannal
func (c *Chan) Eval() Object { return (<-c.ch)() }

// Add ajoute une fonction d'addition de l'objet (Object.Add) à la composition
// de fonction présente pour l'expression
func (c *Chan) Add(obj Object) {
	old := <-c.ch
	c.ch <- func() Object { return old().Add(obj) }
}

// Sub ajoute une fonction de soustraction de l'objet (Object.Sub) à la
// composition de fonction déjà présente pour l'expression
func (c *Chan) Sub(obj Object) {
	old := <-c.ch
	c.ch <- func() Object { return old().Sub(obj) }
}
