// Code generated from Number.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlrp

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NumberLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NumberLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func numberlexerLexerInit() {
	staticData := &NumberLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.SymbolicNames = []string{
		"", "NUMBER",
	}
	staticData.RuleNames = []string{
		"NUMBER", "INT", "EXP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 1, 41, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 3, 0, 9,
		8, 0, 1, 0, 1, 0, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 3, 0, 18, 8,
		0, 1, 0, 3, 0, 21, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1,
		29, 9, 1, 3, 1, 31, 8, 1, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 4, 2, 38, 8,
		2, 11, 2, 12, 2, 39, 0, 0, 3, 1, 1, 3, 0, 5, 0, 1, 0, 4, 1, 0, 48, 57,
		1, 0, 49, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 46, 0, 1, 1,
		0, 0, 0, 1, 8, 1, 0, 0, 0, 3, 30, 1, 0, 0, 0, 5, 32, 1, 0, 0, 0, 7, 9,
		5, 45, 0, 0, 8, 7, 1, 0, 0, 0, 8, 9, 1, 0, 0, 0, 9, 10, 1, 0, 0, 0, 10,
		17, 3, 3, 1, 0, 11, 13, 5, 46, 0, 0, 12, 14, 7, 0, 0, 0, 13, 12, 1, 0,
		0, 0, 14, 15, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 18,
		1, 0, 0, 0, 17, 11, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 20, 1, 0, 0, 0,
		19, 21, 3, 5, 2, 0, 20, 19, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 2, 1, 0,
		0, 0, 22, 31, 5, 48, 0, 0, 23, 27, 7, 1, 0, 0, 24, 26, 7, 0, 0, 0, 25,
		24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0,
		0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 22, 1, 0, 0, 0, 30, 23,
		1, 0, 0, 0, 31, 4, 1, 0, 0, 0, 32, 34, 7, 2, 0, 0, 33, 35, 7, 3, 0, 0,
		34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 37, 1, 0, 0, 0, 36, 38, 7,
		0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39,
		40, 1, 0, 0, 0, 40, 6, 1, 0, 0, 0, 9, 0, 8, 15, 17, 20, 27, 30, 34, 39,
		0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// NumberLexerInit initializes any static state used to implement NumberLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumberLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumberLexerInit() {
	staticData := &NumberLexerLexerStaticData
	staticData.once.Do(numberlexerLexerInit)
}

// NewNumberLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumberLexer(input antlr.CharStream) *NumberLexer {
	NumberLexerInit()
	l := new(NumberLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NumberLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Number.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumberLexerNUMBER is the NumberLexer token.
const NumberLexerNUMBER = 1
