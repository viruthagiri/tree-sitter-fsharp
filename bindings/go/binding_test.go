package tree_sitter_fsharp_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_fsharp "github.com/tree-sitter/tree-sitter-fsharp/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_fsharp.Language())
	if language == nil {
		t.Errorf("Error loading Fsharp grammar")
	}
}

func TestCanLoadSignatureGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_fsharp.LanguageSignature())
	if language == nil {
		t.Errorf("Error loading Fsharp signature grammar")
	}
}
