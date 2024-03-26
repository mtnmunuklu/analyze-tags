package yara

import (
	"bytes"
	"io"

	"github.com/VirusTotal/gyp/ast"
	"github.com/VirusTotal/gyp/parser"
)

func ParseRule(input io.Reader) (rs *ast.RuleSet, err error) {
	return parser.Parse(input)
}

func ParseString(s string) (*ast.RuleSet, error) {
	return ParseRule(bytes.NewBufferString(s))
}

func ParseByte(input []byte) (rs *ast.RuleSet, err error) {
	return ParseRule(bytes.NewBuffer(input))
}
