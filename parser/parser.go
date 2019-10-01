package parser

import (
	"github.com/kgosse/monkeygo/ast"
	"github.com/kgosse/monkeygo/token"

	"github.com/kgosse/monkeygo/lexer"
)

// Parser holds fields of struct
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New returns a new Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram returns the ast root tree
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
