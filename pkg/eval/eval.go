package eval

import (
	"bytes"
	"github.com/Capucinoxx/gox/pkg/object"
	"github.com/Capucinoxx/gox/pkg/token"
)

type Node struct {
	parent *Node
	value  object.Object
	args   []*Node
}

// newNode
func newNode(it *Node, operand object.Object) *Node {
	return &Node{
		parent: it,
		value:  operand,
		args:   make([]*Node, 0),
	}
}

type Eval struct {
	root *Node
	it   *Node
}

func NewEval() *Eval {
	ptr := newNode(nil, nil)
	return &Eval{
		root: ptr,
		it:   ptr,
	}
}

func (e *Eval) Insert(tkn token.Token) {
	// si l'objet que l'on veut ajouter est une parenthèse ouvrante,
	// on descend d'un niveau dans l'arborescence et créant un
	// nouveau noeud. On retourne le poiteur de ce nouveau noeud.
	if tkn.Token() == token.TOKEN_LPAREN {
		ptr := newNode(e.it, nil)
		e.it.args = append(e.it.args, ptr)

		e.it = ptr
		return
	}

	// si l'objet que l'on veut ajouter est une parenthèse fermante, on monte
	// d'un niveau dans l'arborescence. On retourne dle pointeur du parent.
	if tkn.Token() == token.TOKEN_RPAREN {
		// on lance l'ance l'évaluation des enfants
		o := e.it.value.(object.Chan)
		e.it.value = o.Eval()
		e.it.args = nil

		e.it = e.it.parent
		return
	}

	// Sinon l'objet que l'on veut ajouter est un paramètre de la fonction
	// représentant par le pointeur du noeud courant (it) et retourne le pointeur
	// du noeud courant.
	if e.it.value == nil {
		e.it.value = tkn.ToObject()
	} else if _, ok := tkn.ToObject().(object.Null); !ok {
		e.it.value.(object.Chan).Fn(tkn.ToObject())
		e.it.args = append(e.it.args, newNode(e.it, tkn.ToObject()))
	}
}

func (e *Eval) Result() string {
	var out bytes.Buffer

	for i := 0; i < len(e.root.args); i++ {
		out.WriteString(e.root.args[i].value.ToString())
	}

	return out.String()
}
