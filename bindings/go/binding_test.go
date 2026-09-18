package tree_sitter_mdjinja_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_mdjinja "github.com/Benexl/tree-sitter-mdjinja/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_mdjinja.Language())
	if language == nil {
		t.Errorf("Error loading mdjinja grammar")
	}
}
