package lexer

import "github.com/Delta456/interpreter_go/token"

type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	readpos int  // current read position in input (after current char)
	chr     byte // current char under examination
}

// New : inputs a string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

func (l *Lexer) readchar() {
	if l.readpos >= len(l.input) {
		l.chr = 0
	} else {
		l.chr = l.input[l.readpos]
		l.pos = l.readpos
		l.readpos++
	}
}

// NextToken : advances to next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.chr {
	case '=':
		tok = newToken(token.ASSIGN, l.chr)
	case '+':
		tok = newToken(token.PLUS, l.chr)
	//case ':':
	//	tok = newToken(token.COLON, l.chr)
	case ';':
		tok = newToken(token.SEMICOLON, l.chr)
	case '{':
		tok = newToken(token.LBRACE, l.chr)
	case '}':
		tok = newToken(token.RBRACE, l.chr)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.chr) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.chr)
		}
	}

	l.readchar()
	return tok
}

func newToken(tokentype token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokentype, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {

	position := l.pos

	for isLetter(l.chr) {
		l.readchar()
	}

	return l.input[position:l.pos]
}

func isLetter(chr byte) bool {
	return chr >= 'A' && chr <= 'Z' || chr >= 'a' && chr <= 'a' || chr == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.chr == '\n' || l.chr == '\r' || l.chr == '\t' || l.chr == ' ' {
		l.readchar()
	}
}
