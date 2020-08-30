package Parser

import (
	token "github.com/EtoDemerzel0427/acwj/Token"
	"github.com/EtoDemerzel0427/acwj/ast"
	"reflect"
	"strings"
	"testing"
)

func TestPrimary(t *testing.T) {
	tok := token.NewToken(token.Integer, 234)

	n, err := primary(tok)
	if err != nil {
		t.Error(err.Error())
	}
	correct := ast.NewNode(ast.Integer, nil, nil, 234)
	if !reflect.DeepEqual(correct, n) {
		t.Errorf("Wrong value of n: %v.\n", n)
	}
}

func TestParser_BinExpr(t *testing.T) {
	ts := token.NewScanner(strings.NewReader("2"))
	p := NewParser(ts)
	n := p.BinExpr()
	correct := ast.NewLeaf(ast.Integer, 2)
	if !reflect.DeepEqual(correct, n) {
		t.Errorf("Wrong value of n: %v", n)
	}

	ts.Init(strings.NewReader("2 + 151"))
	correct2 := ast.NewLeaf(ast.Integer, 151)
	root := ast.NewNode(ast.Add, correct, correct2, 0)

	n = p.BinExpr()
	if !reflect.DeepEqual(n, root) {
		t.Errorf("Wrong value of n: %v", n)
	} else if !reflect.DeepEqual(n.Left, correct) {
		t.Errorf("Wrong value of n.Left: %v", n.Left)
	} else if !reflect.DeepEqual(n.Right, correct2) {
		t.Errorf("Wrong value of n.Right: %v", n.Right)
	}
}

