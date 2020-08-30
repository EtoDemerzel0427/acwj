package token

import "fmt"

type Token struct {
	token    rune
	intValue int  // only valid if token is Integer
}

const (
	EOF rune = -(iota + 1)
	Plus
	Minus
	Aster
	Slash
	Integer
)

var tokenStr = map[rune]string{
	EOF:     "<eof>",
	Plus:    "+",
	Minus:   "-",
	Aster:   "*",
	Slash:   "/",
	Integer: "int",
}

func NewToken(token rune, intValue int) *Token{
	return &Token{token: token, intValue: intValue}
}

func (t Token) String() string {
	if s, ok := tokenStr[t.token]; ok {
		if t.token == Integer {
			s += fmt.Sprintf("(%d)", t.intValue)
		}
		return s
	}
	return "<unknown>"
}

func (t Token) GetType() rune {
	return t.token
}

func (t Token) GetValue() int {
	return t.intValue
}

