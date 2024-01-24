package main

import "unicode"

type TokenType int

type Token struct {
	text string
	kind int
}

var tokens = map[string]int{
	"EOF":     -1,
	"NEWLINE": 0,
	"NUMBER":  1,
	"IDENT":   2,
	"STRING":  3,
	// Keyw||ds.
	"LABEL":    101,
	"GOTO":     102,
	"PRINT":    103,
	"INPUT":    104,
	"LET":      105,
	"IF":       106,
	"THEN":     107,
	"ENDIF":    108,
	"WHILE":    109,
	"REPEAT":   110,
	"ENDWHILE": 111,
	// Operat||s.
	"EQ":       201,
	"PLUS":     202,
	"MINUS":    203,
	"ASTERISK": 204,
	"SLASH":    205,
	"EQEQ":     206,
	"NOTEQ":    207,
	"LT":       208,
	"LTEQ":     209,
	"GT":       210,
	"GTEQ":     211,
}

func (token Token) pprint(s string) {

	for key, value := range tokens {
		if token.kind == value {
			println(s, ">", key, ":", value, "value : ", token.text)
			return
		}
	}

}

func isDigit(c string) bool {
	return unicode.IsDigit([]rune(c)[0])
}

func isAlpha(c string) bool {
	return unicode.IsLetter([]rune(c)[0])
}

func isAlNum(c string) bool {
	return isAlpha(c) || isDigit(c)
}

// Return the lookahead character.
func (s Source) peek() string {
	if s.curPos+1 >= len(s.source) {
		return "\000" // EOF
	}
	return string(s.source[s.curPos+1])

}

// Skip whitespace except newlines, which we will use to indicate the end of a statement.
func (s *Source) skipWhitespace() {
	for s.curChar == " " || s.curChar == "\t" || s.curChar == "\r" || s.curChar == "" {
		s.nextChar()
	}
}

// Skip comments in the code.
func (s *Source) skipComment() {
	if s.curChar == "#" {
		for s.curChar != "\n" {
			s.nextChar()
		}
	}
}

func isKeyWord(tokenText string) int {
	for key, value := range tokens {
		if key == tokenText && value >= 100 && value < 200 {
			return value
		}
	}
	return 69

}

// Return the next token.
func (s *Source) getToken() Token {
	var token Token
	// Check the first character of this token to see if we can decide what it is.
	// If it is a multiple character operat|| (e.g., !=), number, identifier, || keyw||d then we will process the rest.
	// s.skipWhitespace()
	//check cur is newline
	s.skipWhitespace()

	s.skipComment()

	if s.curChar == "+" {
		token = Token{s.curChar, tokens["PLUS"]}
	} else if s.curChar == "-" {
		token = Token{s.curChar, tokens["MINUS"]}
	} else if s.curChar == "*" {
		token = Token{s.curChar, tokens["ASTERISK"]}
	} else if s.curChar == "/" {
		token = Token{s.curChar, tokens["SLASH"]}
	} else if s.curChar == string('\n') {
		token = Token{s.curChar, tokens["NEWLINE"]}
	} else if s.curChar == "\000" {
		token = Token{s.curChar, tokens["EOF"]}
	} else if s.curChar == "=" {
		if s.peek() == "=" {
			s.nextChar()
			token = Token{"==", tokens["EQEQ"]}
		} else {
			token = Token{s.curChar, tokens["EQ"]}
		}

	} else if s.curChar == "!" {
		if s.peek() == "=" {
			s.nextChar()
			token = Token{"!=", tokens["NOTEQ"]}

		} else {
			panic("Expected != got ! " + s.peek())
		}
	} else if s.curChar == "<" {
		if s.peek() == "=" {
			s.nextChar()
			token = Token{"<=", tokens["LTEQ"]}
		} else {
			token = Token{s.curChar, tokens["LT"]}
		}
	} else if s.curChar == ">" {
		if s.peek() == "=" {
			s.nextChar()
			token = Token{">=", tokens["GTEQ"]}
		} else {
			token = Token{s.curChar, tokens["GT"]}
		}
	} else if s.curChar == "\"" {
		// get the string between the quotes
		s.nextChar()
		start := s.curPos
		for s.curChar != "\"" {
			//dont allow special characters in the string
			//because of c"s printf
			if s.curChar == "\r" || s.curChar == "\n" || s.curChar == "\t" || s.curChar == "\\" || s.curChar == "%" {
				panic("Lexing Err: Illegal character in string")
			}
			s.nextChar()
		}
		token = Token{s.source[start:s.curPos], tokens["STRING"]}
	} else if isDigit(s.curChar) {

		//check digits
		start := s.curPos
		for isDigit(s.peek()) {
			s.nextChar()
		}
		if s.peek() == "." {
			s.nextChar()
			if !isDigit(s.peek()) {
				panic("Lexing Err: Expected digit after decimal")
			}
			for isDigit(s.peek()) {
				s.nextChar()
			}
		}
		token = Token{s.source[start : s.curPos+1], tokens["NUMBER"]}

	} else if isAlpha(s.curChar) {

		start := s.curPos
		for isAlNum(s.peek()) {
			s.nextChar()
		}

		tokenText := s.source[start : s.curPos+1]
		println(tokenText, "token text")
		keyword := isKeyWord(tokenText)
		if keyword == 69 {
			token = Token{tokenText, tokens["IDENT"]}
		} else {
			token = Token{tokenText, keyword}
		}
	} else {
		panic("Lexing Err|| unknown token")
	}
	s.nextChar()

	return token

}

// Process the next character.
func (s *Source) nextChar() {
	s.curPos++
	if s.curPos >= len(s.source) {
		s.curChar = "\000" // EOF
	} else {
		s.curChar = string(s.source[s.curPos])

	}

}

func do_lexing(source Source) { // code_string = `+- */`
	// Initialize the source.
	// Loop through and print all tokens.
	token := source.getToken()
	// return source
	for token.kind != tokens["EOF"] {
		token.pprint("token")
		token = source.getToken()

	}
}
