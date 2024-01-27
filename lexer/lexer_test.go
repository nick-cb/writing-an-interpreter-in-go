package lexer

import (
	"testing"
	"writing-an-interpreter-in-go/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)
	for i, tt := range tests {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Errorf("test [%d] fail. Wrong token type. Expected %q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Errorf("test [%d] fail. Wrong literal value. Expected %q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
