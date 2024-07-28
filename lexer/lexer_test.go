package lexer

import (
	"testing"
	"github.com/aditya-git16/interpreter/token"
)

func TestNextToken(t *testing.T){
	input := `=+(){},;`

	test := []struct{
		expectedType token.TokenType
		expoeectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
}
