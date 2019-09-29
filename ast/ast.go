package ast

import (
	"github.com/kgosse/monkeygo/token"
)

// Node is an interface for the ast tree nodes
type Node interface {
	TokenLiteral() string
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

// Program is the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the TokenLiteral of the first statement
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Identifier is a Node holding an identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the Identifier literal
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement is a Node holding a let statement
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the LetStatement literal
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
