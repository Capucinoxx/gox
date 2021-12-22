package token

const (
	TOKEN_ILLEGAL (uint8) = iota
	TOKEN_EOF
	TOKEN_WS

	/* --- litteral --- */

	TOKEN_IDENTIFIER

	/* --- misc characters --- */

	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_OPERAND
)

type Node struct {
	Parent *Node
	V Token
	args []*Node
}

func NewNode(parent *Node, tkn Token) *Node {
	return &Node{
		Parent: parent,
		V: tkn,
		args: make([]*Node, 0),
	}
}

func (n *Node) Append(el *Node) { n.args = append(n.args, el) }

type Token interface {
	Token() uint8
}


type ILLEGAL struct{ Node }
func (s *ILLEGAL) Token() uint8 { return TOKEN_ILLEGAL }

type EOF struct{ Node }
func (s *EOF) Token() uint8 { return TOKEN_EOF }

type WS struct{ Node }
func (s *WS) Token() uint8  { return TOKEN_WS }

