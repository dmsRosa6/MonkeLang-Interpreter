package parser

import (
	"testing"

	"github.com/dmsRosa/Monke/src/Monke/ast"
	"github.com/dmsRosa/Monke/src/Monke/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
				let x = 5;
				let y = 10;
				let foobar = 838383;`

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != len(tests) {
		t.Fatalf("program.Statements does not contain %d statements. got=%d",
			len(tests), len(program.Statements))
	}

	for i, tt := range tests {
		smtp := program.Statements[i]

		if !testLetStatement(t, smtp, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}
