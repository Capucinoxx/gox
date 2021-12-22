package lexer

import (
	"bufio"
	"bytes"
	"github.com/Capucinoxx/gox/pkg/token"
	"io"
	"strings"
)

var eof = rune(0)

type Scanner struct {
	reader *bufio.Reader
}

// NewScanner retournes une nouvelle instance de Scanner
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(reader)}
}

func (s *Scanner) Scan() (token.Token, string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}

	if isLetter(ch) || isDigit(ch) {
		s.unread()
		return s.scanIdentifier()
	}


	switch ch {
	case eof:
		return &token.EOF{}, ""
	case '(':
		return &token.LPAREN{}, string(ch)
	case ')':
		return &token.RPAREN{}, string(ch)
	case '-', '+', '*', '/':
		if s.read() == '>' {
			return &token.OPERAND{}, string(ch)+ ">"
		}
		s.unread()
		return &token.OPERAND{}, string(ch)

	case '"':
		return s.scanQuote()
	}

	return &token.ILLEGAL{}, string(ch)
}

func (s *Scanner) scanWhitespace() (token.Token, string) {
	// création d'un buffer et fait la lecture du caractère courant à l'intérieur
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// lecture de chaque caractère d'espacement suivant dans le tampon.
	// les caractères qui ne sont pas des caractères d'espacement et eof
	// entraîneront la sortie de la boucle.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return &token.WS{}, buf.String()
}

func (s *Scanner) scanQuote() (token.Token, string) {
	// création d'un buffer et fait la lecture du caractère courant à l'intérieur
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == '"' {
			break
		} else if ch == eof {
			return &token.ILLEGAL{}, "il manque un guillement de fermeture"
		} else {
			buf.WriteRune(ch)
		}
	}

	return &token.IDENTIFIER{}, buf.String()
}

func (s *Scanner) scanIdentifier() (token.Token, string) {
	// création d'un buffer et fait la lecture du caractère courant à l'intérieur
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	str := buf.String()

	switch strings.ToLower(str) {

	}

	return &token.IDENTIFIER{}, str
}

// read lit la prochaine rune du buffered reader.
// Retournes eof (rune(0)) si une erreur se produit (ou si io.EOF est renvoyé)
func (s *Scanner) read() rune {
	ch, _, err := s.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread replaces la rune précédemment lue sur le lecteur
func (s *Scanner) unread() { _ = s.reader.UnreadRune() }

// isWhitespace regarde si le caratère est un espace, une tabluation ou un
// changement de ligne
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
