package lexer

import "github.com/Delta456/interpreter_go/token"

type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	readpos int  // current read position in input (after current char)
	chr     byte // current char under examination
}

// New inputs a string
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
		if l.peekChar() == '=' {
			ch := l.chr
			l.readchar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.chr)}
		} else {
			tok = newToken(token.ASSIGN, l.chr)
		}
	case '+':
		tok = newToken(token.PLUS, l.chr)
	case '-':
		tok = newToken(token.MINUS, l.chr)
	case '!':
		if l.peekChar() == '=' {
			ch := l.chr
			l.readchar()
			tok = token.Token{Type: token.NQ, Literal: string(ch) + string(l.chr)}
		} else {
			tok = newToken(token.BANG, l.chr)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.chr)
	case '>':
		tok = newToken(token.GT, l.chr)
	case '<':
		tok = newToken(token.LT, l.chr)
	case '/':
		tok = newToken(token.SLASH, l.chr)
	case ',':
		tok = newToken(token.COMMA, l.chr)
	case ';':
		tok = newToken(token.SEMICOLON, l.chr)
	case '(':
		tok = newToken(token.LPAREN, l.chr)
	case ')':
		tok = newToken(token.RPAREN, l.chr)
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
		} else if isDigit(l.chr) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
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

func isDigit(chr byte) bool {
	return chr >= '0' && chr <= '9'
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.chr) {
		l.readchar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.readpos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readpos]
	}
}
