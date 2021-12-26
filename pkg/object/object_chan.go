package object

type Chan struct{ ch chan func() Object }

func (c *Chan) Make(tpe Object) {
	c.ch = make(chan func() Object, 1)

	switch tpe.(type) {
	case *Number:
		c.ch <- func() Object { return Number{Value: 0} }
	}

}
func (c *Chan) Add(obj Object) {
	old := <-c.ch
	c.ch <- func() Object { return old().Add(obj) }
}
func (c *Chan) Eval() Object { return (<-c.ch)() }
