package ast

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
	Add:     "+",
	Sub:     "-",
	Mul:     "*",
	Div:     "/",
	Integer: "int",
}

type Node struct {
	Op       rune
	Left     *Node
	Right    *Node
	IntValue int
}

func (n Node) String() string {
	if s, ok := opStr[n.Op]; ok {
		if n.Op == Integer {
			s += fmt.Sprintf("(%d)", n.IntValue)
		}
		if n.Left != nil {
			s += " left: "
			if sl, ok := opStr[n.Left.Op]; ok {
				if n.Left.Op == Integer {
					sl += fmt.Sprintf("(%d)", n.Left.IntValue)
				}
				s += sl
			} else {
				s += "<unknown>"
			}
		}
		if n.Right != nil {
			s += " right: "
			if sr, ok := opStr[n.Right.Op]; ok {
				if n.Right.Op == Integer {
					sr += fmt.Sprintf("(%d)", n.Right.IntValue)
				}
				s += sr
			} else {
				s += "<unknown>"
			}
		}
		s += "\n"


		return s
	}

	return "<unknown>\n"
}

// NewNode build and return a generic AST node
func NewNode(op rune, left *Node, right *Node, intValue int) *Node {
	return &Node{
		Op:       op,
		Left:     left,
		Right:    right,
		IntValue: intValue,
	}
}

// NewLeaf make an AST leaf node
func NewLeaf(op rune, intValue int) *Node {
	return NewNode(op, nil, nil, intValue)
}

// NewUnary make an unary AST node, only one child
func NewUnary(op rune, child *Node, intValue int) *Node {
	return NewNode(op, child, nil, intValue)
}

