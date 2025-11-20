package main_test

import (
	_ "embed"
	"testing"

	"bench/gommep"
	"bench/combp"
	"bench/antlrp"

	"github.com/antlr4-go/antlr/v4"

	"github.com/flowdev/comb"
)

//go:embed 1KB.json
var KB1 string

//go:embed 64KB.json
var KB64 string

//go:embed 512KB.json
var KB512 string

//go:embed 5MB.json
var MB5 string

//go:embed 64KBerr.json
var KB64err string


func Benchmark1KBgomme(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gommep.ParseJSON(KB1)
	}
}

func Benchmark1KBcomb(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = comb.RunOnString(KB1, combp.ParseValue())
	}
}

func Benchmark1KBantlr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := antlr.NewInputStream(KB1)
		lexer := antlrp.NewJSONLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := antlrp.NewJSONParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		_ = p.Json()
	}
}


func Benchmark64KBgomme(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = gommep.ParseJSON(KB64)
	}
}

func Benchmark64KBcomb(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = comb.RunOnString(KB64, combp.ParseValue())
	}
}

func Benchmark64KBantlr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := antlr.NewInputStream(KB64)
		lexer := antlrp.NewJSONLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := antlrp.NewJSONParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		_ = p.Json()
	}
}
