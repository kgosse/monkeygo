package ast

// Node is an interface for the ast tree nodes
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is a kind of Node
type Statement interface {
	Node
	statementNode()
}

// Expression is a kind of Node
type Expression interface {
	Node
	expressionNode()
}
