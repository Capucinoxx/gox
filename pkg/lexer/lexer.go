package lexer

import (
	"bufio"
	"io"
)

var eof = rune(0)

type Scanner struct {
	reader *bufio.Reader
}

// NewScanner retournes une nouvelle instance de Scanner
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(reader)}
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

// skipWhitespace déplace le curseur jusqu'à temps qu'il fasse la lecture
// d'une rune n'étant pas un espace, une tabulation ou un changement de ligne
func (s *Scanner) skipWhitespace(ch rune) rune {
	for isWhitespace(ch) {
		ch = s.read()
	}
	return ch
}

// isWhitespace regarde si le caratère est un espace, une tabluation ou un
// changement de ligne
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
