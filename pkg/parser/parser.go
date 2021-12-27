package parser

import (
	"errors"
	"github.com/Capucinoxx/gox/pkg/eval"
	"github.com/Capucinoxx/gox/pkg/lexer"
	"github.com/Capucinoxx/gox/pkg/token"
	"io"
)

type Parser struct {
	scanner *lexer.Scanner
	eval    *eval.Eval
}

// NewParser retournes une nouvelle instance de Parser
func NewParser(reader io.Reader) *Parser {
	return &Parser{
		scanner: lexer.NewScanner(reader),
		eval:    eval.NewEval(),
	}
}

// Parse fait la lecture du script pour placer les différents éléments
func (p *Parser) Parse() error {
	tkn := p.scan()

	if tkn.Token() != token.TOKEN_LPAREN {
		// TODO gestionnaire erreur position
		return errors.New("")
	}

	for tkn.Token() != token.TOKEN_EOF {
		p.eval.Insert(tkn)
		tkn = p.scan()
	}

	return nil
}

func (p *Parser) scan() token.Token {
	tkn, _ := p.scanner.Scan()
	return tkn
}

func (p *Parser) Result() string { return p.eval.Result() }
