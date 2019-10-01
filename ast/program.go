package ast

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
