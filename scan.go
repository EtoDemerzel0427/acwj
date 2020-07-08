package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
	"unicode"
)

type TokenType int

const (
	T_PLUS TokenType = iota
	T_MINUS
	T_STAR
	T_SLASH
	T_INTLIT
)

type Token struct {
	token TokenType
	intvalue int
}

func next() rune {
	c := TokenScanner.Next()
	if c == '\n' {
		Line++
	}
	return c
}

// skip past input that we don't need to deal with (whitespaces, newlines)
//Return the first character we do need to deal with.
func skip() rune {
	c := next()
	for c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' {
		c = next()
	}
	return c
}

// scan and return an integer literal
func scanint(c rune) int {
	// handle first digit, since c is already scanned.
	k := strings.IndexRune("0123456789", c)  // must >= 0, since isDigit(c) is true
	val := k
	c = TokenScanner.Peek()

	for {
		k = strings.IndexRune("0123456789", c)
		if k < 0 {
			break
		}

		val = val * 10 + k
		TokenScanner.Next()
		c = TokenScanner.Peek()
	}

	return val
}

func (t *Token) Scan() int {
	c := skip()

	switch c {
	case scanner.EOF:
		return 0
	case '+':
		t.token = T_PLUS
	case '-':
		t.token = T_MINUS
	case '*':
		t.token = T_STAR
	case '/':
		t.token = T_SLASH
	default:
		if unicode.IsDigit(c) {
			t.intvalue = scanint(c)
			t.token = T_INTLIT
			break
		}

		fmt.Printf("Unrecognized character %c in line %d\n", c, Line)
		os.Exit(1)
	}

	// found one token
	return 1
}
