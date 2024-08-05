package parser

import (
	"GoInt/ast"
	"GoInt/lexer"
	"GoInt/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func new(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curreToken and peekToken are voth set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
