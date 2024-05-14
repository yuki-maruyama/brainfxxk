package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = "EOF"
	
	MOVR = ">"
	MOVL = "<"
	INCR = "+"
	DECR = "-"

	INPUT = ","
	OUTPUT = "."

	JFOR = "["
	JBAK = "]"
)