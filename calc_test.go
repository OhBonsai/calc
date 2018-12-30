package calc

import (
	"testing"
)

func TestFinal(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		res := Calc(tt.input)
		if res != tt.expected {
			t.Errorf("Wrong answer, got=%d, want=%d", res, tt.expected)
		}
	}
}

func TestTokenizer(t *testing.T) {
	input := `(5 + -10 * 2 + 15 / 3) * 2`
	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{LPAREN, "("},
		{INT, "5"},
		{PLUS, "+"},
		{MINUS, "-"},
		{INT, "10"},
		{ASTERISK, "*"},
		{INT, "2"},
		{PLUS, "+"},
		{INT, "15"},
		{SLASH, "/"},
		{INT, "3"},
		{RPAREN, ")"},
		{ASTERISK, "*"},
		{INT, "2"},
	}

	l := NewLex(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
