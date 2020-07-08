package main

import (
	"strings"
	"testing"
)


func TestSkip(t *testing.T) {
	TokenScanner.Init(strings.NewReader("  \n\r ddd"))
	c := skip()
	if c != 'd' {
		t.Errorf("skip() == %d", c)
	}

	TokenScanner.Init(strings.NewReader("send the letter to me\n"))
	if skip() != 's' {
		t.Error("skip() performs wrong on test 2, case 1")
	}

	TokenScanner.Next()
	TokenScanner.Next()
	TokenScanner.Next()
	TokenScanner.Next()
	if skip() != 't' {
		t.Error("skip() performs wrong on test2 2, case 2")
	}
}

func TestNext(t *testing.T) {
	TokenScanner.Init(strings.NewReader("abcdefghijklmn"))
	c := next()
	if c != 'a' {
		t.Errorf("next() == %c\n", c)
	}
	c = next()
	if c != 'b' {
		t.Errorf("next == %c\n", c)
	}
}

func TestScanint(t *testing.T)  {
	TokenScanner.Init(strings.NewReader("     1234567abcdef"))
	c := skip()
	val := scanint(c)
	if val != 1234567 {
		t.Errorf("Int: val = %d\n", val)
	}
}
