package token

const (
	TOKEN_ILLEGAL (uint8) = iota
	TOKEN_EOF
	TOKEN_WS

	/* --- litteral --- */

	TOKEN_IDENTIFIER

	/* --- misc characters --- */

	TOKEN_SPACE
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_OPERAND

	/* --- keywords --- */

	TOKEN_FILE
)

type token struct {
	Value uint8
	Token Token
}

type Token interface {
	GetToken() uint8
	generate() (string, error)
}

type ILLEGAL struct{ token }

func (s *ILLEGAL) generate() (string, error) { return "", nil }
func (s *ILLEGAL) GetToken() uint8           { return TOKEN_ILLEGAL }

type EOF struct{ token }

func (s *EOF) generate() (string, error) { return "", nil }
func (s *EOF) GetToken() uint8           { return TOKEN_EOF }

type WS struct{ token }

func (s *WS) generate() (string, error) { return "", nil }
func (s *WS) GetToken() uint8           { return TOKEN_WS }

/* --- litteral --- */

type IDENTIFIER struct{ token }

func (s *IDENTIFIER) generate() (string, error) { return "", nil }
func (s *IDENTIFIER) GetToken() uint8           { return TOKEN_IDENTIFIER }

/* --- misc characters --- */

type SPACE struct{ token }

func (s *SPACE) generate() (string, error) { return "", nil }
func (s *SPACE) GetToken() uint8           { return TOKEN_SPACE }

type LPAREN struct{ token }

func (s *LPAREN) generate() (string, error) { return "", nil }
func (s *LPAREN) GetToken() uint8           { return TOKEN_LPAREN }

type RPAREN struct{ token }

func (s *RPAREN) generate() (string, error) { return "", nil }
func (s *RPAREN) GetToken() uint8           { return TOKEN_RPAREN }

type OPERAND struct{ token }

func (s *OPERAND) generate() (string, error) { return "", nil }
func (s *OPERAND) GetToken() uint8           { return TOKEN_OPERAND }

/* --- keywords --- */

type FILE struct{ token }

func (s *FILE) generate() (string, error) { return "", nil }
func (s *FILE) GetToken() uint8           { return TOKEN_FILE }
