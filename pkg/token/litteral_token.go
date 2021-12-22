package token

type IDENTIFIER struct{ Node }

func (s *IDENTIFIER) Token() uint8 { return TOKEN_IDENTIFIER }
