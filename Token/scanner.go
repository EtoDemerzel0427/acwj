package token

import (
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

// TODO: add more members
type TokenScanner struct {
	Tok *Token
	Scanner
}

// Encapsulate scanner.Init
func (ts *TokenScanner) init(reader io.Reader) *TokenScanner {
	ts.Init(reader)

	return ts
}

func NewScanner(reader io.Reader) *TokenScanner {
	ts := &TokenScanner{
		Tok: &Token{0, 0}, // 0 indicates no token, not even EOF
	}

	return ts.init(reader)
}


// skip past input that we don't need to deal with (i.e, whitespaces)
// Return the first character we do need to deal with.
func (ts *TokenScanner) skip() rune {
	c := ts.Next()
	for c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' {
		c = ts.Next()
	}
	return c
}

// TODO: refer scanner/scanner scanNumber
// scan and return an integer literal
func (ts *TokenScanner) scanInt(c rune) int {
	// handle first digit, since c is already scanned.
	k := strings.IndexRune("0123456789", c)  // must >= 0, since isDigit(c) is true
	val := k
	c = ts.Peek()

	for {
		k = strings.IndexRune("0123456789", c)
		if k < 0 {
			break
		}

		val = val * 10 + k
		ts.Next()
		c = ts.Peek()
	}

	return val
}


func (ts *TokenScanner) Scan() (int, error) {
	c := ts.skip() // go to the first non-whitespace char
	switch c {
	case scanner.EOF:
		ts.Tok = &Token{token: EOF}
		return 0, nil  // no token
	case '+':
		ts.Tok = &Token{token: Plus}
	case '-':
		ts.Tok = &Token{token: Minus}
	case '*':
		ts.Tok = &Token{token: Aster}
	case '/':
		ts.Tok = &Token{token: Slash}
	default:
		if unicode.IsDigit(c) {
			ts.Tok = &Token{token: Integer, intValue: ts.scanInt(c)}
			break
		}

		err := fmt.Errorf("Unrecognized character: %c.\n", c)
		return 0, err
	}
	return 1, nil  // found one token
}



