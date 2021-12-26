package parser

import (
	"errors"
	"github.com/Capucinoxx/gox/pkg/lexer"
	"github.com/Capucinoxx/gox/pkg/token"
	"io"
)

type Parser struct {
	scanner *lexer.Scanner
	root    *token.Node
}

// NewParser retournes une nouvelle instance de Parser
func NewParser(reader io.Reader) *Parser {
	return &Parser{
		scanner: lexer.NewScanner(reader),
		root:    token.NewNode(nil, nil),
	}
}

// insert insère un nouvel objet de type token.Token dans une structure
// en arborescence représentant la structure du script venant d'être
// parser
func (p *Parser) insert(it *token.Node, tkn token.Token) *token.Node {
	// si l'objet que l'on veut ajouter est une parenthèse ouvrante,
	// on descend d'un niveau dans l'arborescence et créant un
	// nouveau noeud. On retourne le poiteur de ce nouveau noeud.
	if tkn.Token() == token.TOKEN_LPAREN {
		ptr := token.NewNode(it, tkn)
		it.Append(ptr)

		return ptr
	}

	// si l'objet que l'on veut ajouter est une parenthèse fermante, on monte
	// d'un niveau dans l'arborescence. On retourne dle pointeur du parent.
	if tkn.Token() == token.TOKEN_RPAREN {
		return it.Parent
	}

	// Sinon l'objet que l'on veut ajouter est un paramètre de la fonction
	// représentant par le pointeur du noeud courant (it) et retourne le pointeur
	// du noeud courant.

	if it.V.Token() == token.TOKEN_LPAREN {
		it.V = tkn
	} else {
		it.Append(token.NewNode(it.Parent, tkn))
	}

	return it
}

// Parse fait la lecture du script pour placer les différents éléments
func (p *Parser) Parse() error {
	tkn := p.scan()

	if tkn.Token() != token.TOKEN_LPAREN {
		// TODO gestionnaire erreur position
		return errors.New("")
	}

	it := p.root
	for tkn.Token() != token.TOKEN_EOF {
		it = p.insert(it, tkn)
		tkn = p.scan()
	}

	return nil
}

func (p *Parser) scan() token.Token {
	tkn, _ := p.scanner.Scan()
	return tkn
}

func (p *Parser) CallStack() string {
	return token.CallStack(p.root.FirstArgs())
}
