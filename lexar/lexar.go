package lexar

import "brainfxxk/token"

type Lexar struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexar {
	l := &Lexar{input: input}
	l.readChar()
	return l
}

func (l *Lexar) readChar(){
	if l.readPosition >= len(l.input){
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexar) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '<':
		tok = newToken(token.MOVL, l.ch)
	case '>':
		tok = newToken(token.MOVR, l.ch)
	case '+':
		tok = newToken(token.INCR, l.ch)
	case '-':
		tok = newToken(token.DECR, l.ch)
	case ',':
		tok = newToken(token.INPUT, l.ch)
	case '.':
		tok = newToken(token.OUTPUT, l.ch)
	case '[':
		tok = newToken(token.JFOR, l.ch)
	case ']':
		tok = newToken(token.JBAK, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
} 

func (l *Lexar) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}