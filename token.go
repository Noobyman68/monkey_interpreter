package token

type token_type string

type token struct {
  Type token_type
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENT = "IDENT"
  INT = "INT"

  ASSIGN = "=" 
  PLUS = "+" 

  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  FUNCTION = "FUNCTION"
  LET = "LET"
)

