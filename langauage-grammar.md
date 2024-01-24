program ::= {statement}
statement ::= "DEKHAO" (expression | string) nl
| "AGR" comparison "PHR" nl {statement} "AGRBND" nl
| "JAB" comparison "KARO" nl {statement} "JABBND" nl
| "YE" ident nl
| "JAO" ident nl
| "NAM" ident "=" expression nl
| "BTAO" ident nl
comparison ::= expression (("==" | "!=" | ">" | ">=" | "<" | "<=") expression)+
expression ::= term {( "-" | "+" ) term}
term ::= unary {( "/" | "\*" ) unary}
unary ::= ["+" | "-"] primary
primary ::= number | ident
nl ::= '\n'

program ::= {statement}
statement ::= "PRINT" (expression | string) nl
| "IF" comparison "THEN" nl {statement} "ENDIF" nl
| "WHILE" comparison "REPEAT" nl {statement} "ENDWHILE" nl
| "LABEL" ident nl
| "GOTO" ident nl
| "LET" ident "=" expression nl
| "INPUT" ident nl
