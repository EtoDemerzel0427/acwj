package main

type NodeType int

const (
	A_ADD NodeType = iota
	A_SUBTRACT
	A_MULTIPLY
	A_DIVIDE
	A_INITLIT
)

type ASTNode struct {
	op NodeType
	left *ASTNode
	right *ASTNode
	intvalue int
}
