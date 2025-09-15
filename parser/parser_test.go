package parser

import (
	"monkey-language/ast"
	"monkey-language/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      int64
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", 1},
		{"let foobar = y;", "foobar", 2},
		// {"let 213", "", 0}, // for check errors
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier, tt.expectedValue) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string, value int64) bool {
	if s.TokenLiteral() != "let" {
		t.Fatalf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	if letStmt, ok := s.(*ast.LetStatement); ok {
		if letStmt.Name.Value != name {
			t.Fatalf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
			return false
		}
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return 123456;`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got = %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement. got = %T", stmt)
			continue
		}

		if retStmt.TokenLiteral() != "return" {
			t.Errorf("retStmt.TokenLiteral is not 'return', got = %q", retStmt.TokenLiteral())
		}
	}
}
