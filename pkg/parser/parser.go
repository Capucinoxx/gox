package parser

import (
	"github.com/Capucinoxx/gox/pkg/lexer"
	"github.com/Capucinoxx/gox/pkg/token"
	"io"
)

type Parser struct {
	scanner *lexer.Scanner
	buffer  struct {
		tkn     token.Token // dernier token lue
		literal string      // dernière description lue
		size    uint64      // taille du tampon
	}
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{scanner: lexer.NewScanner(reader)}
}

func (p *Parser) Parse() (string, error) {
	return "", nil
}

func (p *Parser) scan() (token.Token, string) {
	if p.buffer.size != 0 {
		p.buffer.size = 0
		return p.buffer.tkn, p.buffer.literal
	}

	tkn, lit := p.scanner.Scan()

	p.buffer.tkn, p.buffer.literal = tkn, lit

	return tkn, lit
}

func (p *Parser) unscan() { p.buffer.size = 1 }

func (p *Parser) scanIgnoreWhitespace() (token.Token, string) {
	tkn, literal := p.scan()
	if tkn == token.WS {
		tkn, literal = p.scan()
	}
	return tkn, literal
}
