package parser

import (
	"fmt"

	"github.com/yuki-maruyama/brainfxxk/ast"
	"github.com/yuki-maruyama/brainfxxk/lexar"
	"github.com/yuki-maruyama/brainfxxk/token"
)

type Parser struct {
	l *lexar.Lexar

	curToken token.Token
	peekToken token.Token
}

func Parse(s string) (*ast.Program, error) {
	l := lexar.New(s)
	p := New(l)
	return p.ParseProgram()
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

func (p *Parser) ParseProgram() (*ast.Program, error) {
	var parse func() ([]ast.Node, error)
	var stack int = 1
	parse = func() ([]ast.Node, error) {
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
				p.NextToken()
				loopBody, err := parse()
				if err != nil{
					return nil, err
				}
				body = append(body, &ast.Loop{Body: loopBody})
				stack++
			case token.JBAK:
				if body == nil {
					return nil, fmt.Errorf("empty while block")
				}
				stack--
				if stack <= 1 {
					return nil, fmt.Errorf("unexpected token ']'")
				}
				return body, nil
			}
			p.NextToken()
		}

		if stack != 1 {
			return nil, fmt.Errorf("unclosed while block")
		}
		return body, nil
	}
	
	body, err := parse()
	if err != nil{
		return nil, err
	}
	return &ast.Program{Body: body}, nil
}