package parser

import (
	"testing"
)

func TestParseLineSuccess(t *testing.T) {
	input := "commands/check.go:10:1: exported function Check should have comment or be unexported"
	annotation, err := ParseLine(input)

	if err != nil {
		t.Error(err)
	}

	if annotation.Path != "commands/check.go" {
		t.Errorf("Path: got %q, want %q", annotation.Path, "commands/check.go")
	}

	if *annotation.StartLine != 10 {
		t.Errorf("StartLine: got %d, want %d", *annotation.StartLine, 10)
	}

	if *annotation.StartColumn != 1 {
		t.Errorf("StartColumn: got %d, want %d", *annotation.StartColumn, 1)
	}
}

func TestParseLineFail(t *testing.T) {
	if _, err := ParseLine(""); err == nil {
		t.Errorf("Expeced error for blank string, got none")
	}

	if _, err := ParseLine("wat"); err == nil {
		t.Errorf("Expeced error for nonsense string, got none")
	}

	if _, err := ParseLine("bad/line/number:bar:1: msg"); err == nil {
		t.Errorf("Expeced failure parsing line number, got none")
	}

	if _, err := ParseLine("bad/column/number:1:nope: msg"); err == nil {
		t.Errorf("Expeced failure parsing column number, got none")
	}
}
