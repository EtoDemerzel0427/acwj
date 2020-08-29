package Parser

import "fmt"

const (
	EOF = -(iota + 1)
	Add
	Sub
	Mul
	Div
	Integer
)

// todo: integrate ast with token
var opStr = map[rune]string{
	EOF:     "<eof>",
	Add:    "+",
	Sub:   "-",
	Mul:   "*",
	Div:   "/",
	Integer: "int",
}

type Node struct {
	op rune
	left *Node
	right *Node
	intValue int
}

func (n Node) String() string {
	if s, ok := opStr[n.op]; ok {
		if n.op == Integer {
			s += fmt.Sprintf("(%d)", n.intValue)
		}
		return s
	}

	return "<unknown>"
}

// build and return a generic AST node
func NewNode(op rune, left *Node, right *Node, intValue int) *Node {
	return &Node{
		op: op,
		left: left,
		right: right,
		intValue: intValue,
	}
}

// make an AST leaf node
func NewLeaf(op rune, intValue int) *Node {
	return NewNode(op, nil, nil, intValue)
}

// make an unary AST node, only one child
func NewUnary(op rune, child *Node, intValue int) *Node {
	return NewNode(op, child, nil, intValue)
}

