package Parser

import (
	"errors"
	"fmt"
	token "github.com/EtoDemerzel0427/acwj/Token"
	"github.com/EtoDemerzel0427/acwj/ast"
	"log"
)

var tokMap = map[rune]rune{
	token.EOF:   ast.EOF,
	token.Plus:  ast.Add,
	token.Minus: ast.Sub,
	token.Aster: ast.Mul,
	token.Slash: ast.Div,
}

var opPrec = map[rune]int{
	token.EOF:     0,
	token.Plus:    10,
	token.Minus:   10,
	token.Aster:   20,
	token.Slash:   20,
	token.Integer: 0,
}

type Parser struct {
	ts *token.TokenScanner
}

func OpPrecedence(t rune) (int, error) {
	if prec, ok := opPrec[t]; ok && prec != 0 {
		return prec, nil
	} else {
		return -1, errors.New("Syntax error: Unrecognized op")
	}
}

func NewParser(ts *token.TokenScanner) *Parser {
	return &Parser{
		ts: ts,
	}
}

func primary(t *token.Token) (*ast.Node, error) {
	switch t.GetType() {
	case token.Integer:
		return ast.NewLeaf(ast.Integer, t.GetValue()), nil
	default:
		return nil, errors.New("Syntax error: Unexpected token type " + t.String())
	}
}

func (p *Parser) next() {
	_, err := p.ts.Scan()
	if err != nil {
		log.Fatal(p.ts.Pos().String() + ": " + err.Error())
	}
}

func (p *Parser) BinExpr(prec int) *ast.Node {
	p.next()
	left, err := primary(p.ts.Tok)
	if err != nil {
		log.Fatal(p.ts.Pos().String() + ": " + err.Error())
	}

	p.next()
	t := p.ts.Tok.GetType()
	if t == token.EOF {
		return left
	}

	for  {
		nextPrec, err := OpPrecedence(t)
		if err != nil {
			log.Fatal(p.ts.Pos().String() + ": " + err.Error())
		}

		if nextPrec <= prec {
			break
		}

		// p.next()
		right := p.BinExpr(nextPrec)

		left = ast.NewNode(tok2op(t), left, right, 0)

		t = p.ts.Tok.GetType()  // update t
		if t == token.EOF {
			return left
		}
	}



	return left
}

func tok2op(t rune) rune {
	if op, ok := tokMap[t]; ok {
		return op
	}

	return 0 // valid token ranges from -6 to -1
}

func InterpretTree(n *ast.Node) (int, error) {
	var leftVal, rightVal int
	var err error
	if n.Left != nil {
		leftVal, err = InterpretTree(n.Left)
		if err != nil {
			return 0, errors.New(err.Error())
		}
	}
	if n.Right != nil {
		rightVal, err = InterpretTree(n.Right)
		if err != nil {
			return 0, errors.New(err.Error())
		}
	}

	fmt.Printf("%v", n)

	switch n.Op {
	case ast.Add:
		return leftVal + rightVal, nil
	case ast.Sub:
		return leftVal - rightVal, nil
	case ast.Mul:
		return leftVal * rightVal, nil
	case ast.Div:
		return leftVal / rightVal, nil
	case ast.Integer:
		return n.IntValue, nil
	default:
		return 0, errors.New("Unknown AST node.\n")

	}
}
