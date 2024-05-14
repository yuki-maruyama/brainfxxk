package parser

import (
	"brainfxxk/ast"
	"brainfxxk/lexar"
	"brainfxxk/token"
)

type Parser struct {
	l *lexar.Lexar

	curToken token.Token
	peekToken token.Token
}

func New(l *lexar.Lexar) *Parser {
	p := &Parser{l: l}

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	var parse func() []ast.Node
	parse = func() []ast.Node {
		var body []ast.Node
		for p.curToken.Type != token.EOF {
			switch p.curToken.Type{
			case token.MOVR:
				body = append(body, &ast.MoveRight{})
			case token.MOVL:
				body = append(body, &ast.MoveLeft{})
			case token.INCR:
				body = append(body, &ast.Increment{})
			case token.DECR:
				body = append(body, &ast.Decrement{})
			case token.INPUT:
				body = append(body, &ast.Input{})
			case token.OUTPUT:
				body = append(body, &ast.Output{})
			case token.JFOR:
				body = append(body, &ast.Loop{Body: parse()})
			case token.JBAK:
				return body
			}
			p.NextToken()
		}
		return body
	}
	
	return &ast.Program{Body: parse()}
}