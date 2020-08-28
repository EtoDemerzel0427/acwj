package main

import (
	"fmt"
	"log"
	"os"
)

type NodeType int

const (
	AddNode NodeType = iota
	SubtractNode
	MULTIPLYNode
	DIVIDENode
	INTLITNode
)

var astOp = [...]string{
	AddNode: "+",
	SubtractNode: "-",
	MULTIPLYNode: "*",
	DIVIDENode: "/",
	INTLITNode: "int_literal"}

type ASTNode struct {
	op NodeType
	left *ASTNode
	right *ASTNode
	intValue int
}

func (n *ASTNode) String() string {
	return astOp[n.op]
}

func (n *ASTNode) InterpretAST() int {
	leftVal, rightVal := 0, 0
	if n.left != nil {
		leftVal = n.left.InterpretAST()
	}
	if n.right != nil {
		rightVal = n.right.InterpretAST()
	}

	// if n is INTLIT, no leftVal/rightVal
	if n.op == INTLITNode {
		fmt.Printf("int %d\n", n.intValue)
	} else {
		fmt.Printf("%d %s %d\n", leftVal, n.String(), rightVal)
	}

	switch n.op {
	case AddNode:
		return leftVal + rightVal
	case SubtractNode:
		return leftVal - rightVal
	case MULTIPLYNode:
		return leftVal * rightVal
	case DIVIDENode:
		return leftVal / rightVal
	case INTLITNode:
		return n.intValue
	default:
		log.Fatalf("Unknown AST op %d\n", n.op)
		os.Exit(-1)
		return -1  // todo: deal with error in a beautiful manner
	}
}

// build and return a generic AST node
func NewASTNode(op NodeType, left *ASTNode, right *ASTNode, intValue int) *ASTNode {
	return &ASTNode{
		op: op,
		left: left,
		right: right,
		intValue: intValue,
	}
}

// make an AST leaf node
func NewASTLeaf(op NodeType, intValue int) *ASTNode {
	return NewASTNode(op, nil, nil, intValue)
}

// make an unary AST node, only one child
func NewASTUnary(op NodeType, child *ASTNode, intValue int) *ASTNode {
	return NewASTNode(op, child, nil, intValue)
}

// parse a primary factor and return its ASTNode representation
func Primary() *ASTNode {
	switch CurToken.token {
	case IntLitToken:
		n := NewASTLeaf(INTLITNode, CurToken.intvalue)
		CurToken.Scan()  // fetch next to CurToken
		return n
	default:
		log.Fatalf("syntax error in line %d\n", Line)
		return nil
	}
}

func arithop(tok TokenType) NodeType {
	switch tok {
	case PlusToken:
		return AddNode
	case MinusToken:
		return SubtractNode
	case StarToken:
		return MULTIPLYNode
	case SlashToken:
		return DIVIDENode
	default:
		log.Fatalf("Unknown token in arithop() in line %d\n", Line)
		return -1
	}
}

func Binexpr() *ASTNode {
	left := Primary()
	if CurToken.token == EOFToken {
		return left
	}

	nodetype := arithop(CurToken.token)

	CurToken.Scan()

	right := Binexpr()

	return NewASTNode(nodetype, left, right, 0)
}






