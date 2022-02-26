package token

type TokenType string

type Token struct{
	Type TokenType
	Literal String
}

const{
	ILL = "ILLEGAL"
	EOF = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT = "INT"

	EQ = "="
	PLUS = "+"
	MINUS = "-"

	COMMA = ","
	SEMICOLON = ";"
	LP = "("
	RP = ")"
	LC = "{"
	RC = "}"

	FUNC = "FUNCTION"
	LET = "LET"

}