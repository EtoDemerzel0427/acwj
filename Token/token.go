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

func (t Token) String() string {
	if s, ok := tokenStr[t.token]; ok {
		if t.token == Integer {
			s += fmt.Sprintf("(%d)", t.intValue)
		}
		return s
	}
	return "<unknown>"
}


