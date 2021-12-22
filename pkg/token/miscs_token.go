package token

type LPAREN struct{ Node }

func (s *LPAREN) Token() uint8 { return TOKEN_LPAREN }

type RPAREN struct{ Node }

func (s *RPAREN) Token() uint8 { return TOKEN_RPAREN }

type OPERAND struct{ Node }

func (s *OPERAND) Token() uint8 { return TOKEN_OPERAND }
