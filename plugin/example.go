// This must be package main
package main

import (
	linters "github.com/dbraley/example-linter"
	"golang.org/x/tools/go/analysis"
)

// This must be implemented
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{linters.TodoAnalyzer}, nil
}

