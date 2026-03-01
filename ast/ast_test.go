package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	expected := "let myVar = anotherVar;"

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != expected {
		t.Errorf("program.String() wrong. expected=%q, got=%q", expected, program.String())
	}
}

func TestStringMultiple(t *testing.T) {
	expected := "let myVar = anotherVar;return 10;"

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "10"},
					Value: "10",
				},
			},
		},
	}

	if program.String() != expected {
		t.Errorf("program.String() wrong. expected=%q, got=%q", expected, program.String())
	}
}
