package main

// import "C"

type Parser struct {
	source    Source
	curToken  Token
	peekToken Token
	emitter   Emitter
}

// Return true if the current token matches.
func (p Parser) checkToken(kind int) bool {
	return kind == p.curToken.kind
}

// Return true if the next token matches.
func (p Parser) checkPeek(kind int) bool {
	return kind == p.peekToken.kind
}
func getTokenText(kind int) string {
	for key, value := range tokens {
		if kind == value {
			return key
		}
	}
	return ""
}

// Try to match current token. If not, error. Advances the current token.
func (p *Parser) match(kind int) {
	if !p.checkToken(kind) {
		msg := "Expected " + getTokenText(kind) + " got " + getTokenText(p.curToken.kind)
		panic(msg)
	}
	p.nextToken()
}

// Advances the current token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.source.getToken()
}

//   ******************************************* PRODUCTION RULES *******************************************

func (p *Parser) nl() {
	println("NEW LINE")
	//require atleast one newline
	p.match(tokens["NEWLINE"])
	//skip any extra newlines
	for p.checkToken(tokens["NEWLINE"]) {
		p.nextToken()
	}

}

// primary ::= number | ident

func (p *Parser) primary() {
	println("PRIMARY(", p.curToken.text, ")")
	if p.checkToken(tokens["NUMBER"]) {
		p.emitter.emit(p.curToken.text)
		p.nextToken()
	} else if p.checkToken(tokens["IDENT"]) {
		//ensure the variable already exists
		if !p.source.symbols.Contains(p.curToken.text) {
			msg := "Referencing variable before assignment -> " + p.curToken.text
			panic(msg)
		}
		p.emitter.emit(p.curToken.text)

		p.nextToken()
	} else {
		msg := "Unexpected token at start of primary: " + getTokenText(p.curToken.kind)
		panic(msg)
	}
}

// primary ::= number | ident

func (p *Parser) unary() {
	println("UNARY")
	if p.checkToken(tokens["MINUS"]) || p.checkToken(tokens["PLUS"]) {
		p.emitter.emit(p.curToken.text)

		p.nextToken()
	}
	p.primary()
}

// term ::= unary {( "/" | "*" ) unary}

func (p *Parser) term() {
	println("TERM")
	p.unary()
	// # Can have 0 or more *// and expressions.

	for p.checkToken(tokens["ASTERISK"]) || p.checkToken(tokens["SLASH"]) {
		p.emitter.emit(p.curToken.text)

		p.nextToken()
		p.unary()

	}

}

// expression ::= term {( "-" | "+" ) term}

func (p *Parser) expression() {
	println("EXPRESSION")
	p.term()
	for p.checkToken(tokens["PLUS"]) || p.checkToken(tokens["MINUS"]) {
		p.emitter.emit(p.curToken.text)

		p.nextToken()
		p.term()
	}

}

func (p *Parser) isComparisonOperator() bool {
	return p.checkToken(tokens["EQEQ"]) || p.checkToken(tokens["NOTEQ"]) || p.checkToken(tokens["GT"]) || p.checkToken(tokens["GTEQ"]) || p.checkToken(tokens["LT"]) || p.checkToken(tokens["LTEQ"])
}

// comparison ::= expression (("==" | "!=" | ">" | ">=" | "<" | "<=") expression)+

func (p *Parser) comparison() {
	println("comparison")

	p.expression()
	// Must be at least one comparison operator and another expression.
	if p.isComparisonOperator() {
		p.emitter.emit(p.curToken.text)

		p.nextToken()
		p.expression()
	} else {
		msg := "Expected comparison operator, got " + getTokenText(p.curToken.kind)
		panic(msg)
	}
	for p.isComparisonOperator() {
		p.emitter.emit(p.curToken.text)

		p.nextToken()
		p.expression()
	}

}

// statements
func (p *Parser) statement() {
	//Check the first token to see what kind of statement this is.
	// # "PRINT" (expression | string)

	if p.checkToken(tokens["PRINT"]) {
		// "PRINT" (expression | string)

		println("PRINT-statement")

		p.nextToken()
		if p.checkToken(tokens["STRING"]) {
			//string
			println("STRING", p.curToken.text)
			pr := "printf(\"" + p.curToken.text + "\\n\"" + ");"
			p.emitter.emitLine(pr)

			p.nextToken()
		} else {
			//expression
			p.emitter.emit("printf(\"%" + ".2f\\n\", (float)(")
			p.expression()
			p.emitter.emitLine("));")

		}

	} else if p.checkToken((tokens["IF"])) {
		// "IF" comparison "THEN" {statement} "ENDIF"
		println("IF-statement")
		p.nextToken()
		p.emitter.emit("if(")

		p.comparison()
		p.match(tokens["THEN"])
		p.nl()
		p.emitter.emitLine("){")

		for !p.checkToken(tokens["ENDIF"]) {
			p.statement()
		}
		p.match(tokens["ENDIF"])
		p.emitter.emitLine("}")

	} else if p.checkToken((tokens["WHILE"])) {
		// "WHILE" comparison "REPEAT" {statement} "ENDWHILE"
		println("WHILE-statement")
		p.nextToken()
		p.emitter.emit("while(")

		p.comparison()
		p.match(tokens["REPEAT"])
		p.nl()
		p.emitter.emitLine("){")

		for !p.checkToken(tokens["ENDWHILE"]) {
			p.statement()
		}
		p.match(tokens["ENDWHILE"])
		p.emitter.emitLine("}")

	} else if p.checkToken((tokens["LABEL"])) {
		// "LABEL" ident
		println("LABEL-statement")
		p.nextToken()
		// Make sure this label doesn't already exist.
		if p.source.labelsDeclared.Contains(p.curToken.text) {
			msg := "Label already exists " + p.curToken.text
			panic(msg)
		}
		p.source.labelsDeclared.Add(p.curToken.text)
		p.emitter.emitLine(p.curToken.text + ":")

		p.match(tokens["IDENT"])
	} else if p.checkToken((tokens["GOTO"])) {
		// "GOTO" ident
		println("GOTO-statement")
		p.nextToken()
		p.source.labelsGotoed.Add(p.curToken.text)
		p.emitter.emitLine("goto " + p.curToken.text + ";")

		p.match(tokens["IDENT"])
	} else if p.checkToken((tokens["LET"])) {
		// "LET" ident "=" expression
		println("LET-statement")
		p.nextToken()
		//            #  Check if ident exists in symbol table. If not, declare it.

		if !p.source.symbols.Contains(p.curToken.text) {
			p.source.symbols.Add(p.curToken.text)
			p.emitter.headerLine("float " + p.curToken.text + ";")
		}
		p.emitter.emit(p.curToken.text + " = ")
		p.match(tokens["IDENT"])
		p.match(tokens["EQ"])
		p.expression()
		p.emitter.emitLine(";")

	} else if p.checkToken((tokens["INPUT"])) {
		// "INPUT" ident
		println("INPUT-statement")
		p.nextToken()
		//If variable doesn't already exist, declare it.
		if !p.source.symbols.Contains(p.curToken.text) {
			p.source.symbols.Add(p.curToken.text)
			p.emitter.headerLine("float " + p.curToken.text + ";")

		}
		p.emitter.emitLine("if(0 == scanf(\"%" + "f\", &" + p.curToken.text + ")) {")
		p.emitter.emitLine(p.curToken.text + " = 0;")
		p.emitter.emit("scanf(\"%")
		p.emitter.emitLine("*s\");")
		p.emitter.emitLine("}")
		p.match(tokens["IDENT"])
	} else {
		//error
		println("ERROR")
		msg := "Unexpected token at start of statement: " + getTokenText(p.curToken.kind) + "(" + p.curToken.text + ")"
		panic(msg)
	}

	p.nl()

}

// program ::= {statement}
func (p *Parser) program() {
	println("program")
	//skip any newlines at the start
	p.emitter.headerLine("#include <stdio.h>")
	p.emitter.headerLine("int main(void){")
	for p.checkToken(tokens["NEWLINE"]) {
		p.nextToken()
	}

	for !p.checkToken(tokens["EOF"]) {
		p.statement()
	}
	p.emitter.emitLine("return 0;")
	p.emitter.emitLine("}")
	// Check that each label referenced in a GOTO is declared.
	for key, _ := range p.source.labelsGotoed.m {
		if !p.source.labelsDeclared.Contains(key) {
			msg := "Attempting to GOTO to undeclared label: " + key
			panic(msg)
		}
	}

}

// ************************* MAIN  *************************
func do_parsing(p Parser) {
	//called two times to set curToken and peekToken
	p.nextToken()
	p.nextToken()

	p.program()
	p.emitter.writeFile()

	println("DONE PARSING")

}
