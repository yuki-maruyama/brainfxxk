// lexar/lexar_test.go

package lexar

import (
    "brainfxxk/token"
    "testing"
)

func TestNextToken(t *testing.T) {
    input := `<>+-,.[]`
    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.MOVL, "<"},
        {token.MOVR, ">"},
        {token.INCR, "+"},
        {token.DECR, "-"},
        {token.INPUT, ","},
        {token.OUTPUT, "."},
        {token.JFOR, "["},
        {token.JBAK, "]"},
        {token.EOF, ""},
    }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }
}

func TestSkipWhiteSpace(t *testing.T) {
    input := `  +   -`
    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.INCR, "+"},
        {token.DECR, "-"},
        {token.EOF, ""},
    }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }
}
