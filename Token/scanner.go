package token

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/scanner"
	"unicode"
)

// TODO: add more members
type tokenScanner struct {
	tok Token
	scanner.Scanner
}

// Encapsulate scanner.Init
func (ts *tokenScanner) Init(reader io.Reader) *tokenScanner {
	ts.Init(reader)

	return ts
}

func NewScanner(reader io.Reader) *tokenScanner {
	tsc := &tokenScanner{
		tok: Token{-2, 0},
	}

	return tsc.Init(reader)
}

// todo: add error handling in other func
func (ts *tokenScanner) error(msg string) {
	ts.ErrorCount++

	pos := ts.Position
	if !pos.IsValid() {
		pos = ts.Pos()
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", pos, msg)
}

func (ts *tokenScanner) errorf(format string, args ...interface{}) {
	ts.error(fmt.Sprintf(format, args...))
}

// skip past input that we don't need to deal with (i.e, whitespaces)
// Return the first character we do need to deal with.
func (ts *tokenScanner) skip() rune {
	c := ts.Next()
	for c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' {
		c = ts.Next()
	}
	return c
}

// TODO: refer scanner/scanner scanNumber
// scan and return an integer literal
func (ts *tokenScanner) scanInt(c rune) int {
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


func (ts *tokenScanner) Scan() int {
	c := ts.skip() // go to the first non-whitespace char
	switch c {
	case scanner.EOF:
		ts.tok.token = EOF
		return 0
	case '+':
		ts.tok.token = Plus
	case '-':
		ts.tok.token = Minus
	case '*':
		ts.tok.token = Aster
	case '/':
		ts.tok.token = Slash
	default:
		if unicode.IsDigit(c) {
			ts.tok.token = Integer
			ts.tok.intValue = ts.scanInt(c)
			break
		}

		ts.error("Unrecognized character.")
	}
	return 0
}



