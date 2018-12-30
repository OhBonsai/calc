package calc

import "github.com/OhBonsai/yolang/token"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"

	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LPAREN = "("
	RPAREN = ")"
)

func Calc(input string) int64 {

}

type Token struct {
	Type    string
	Literal string
}

func newToken(tokenType string, c byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(c),
	}
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLex(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
