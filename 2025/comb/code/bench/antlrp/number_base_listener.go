// Code generated from Number.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlrp // Number
import "github.com/antlr4-go/antlr/v4"

// BaseNumberListener is a complete listener for a parse tree produced by NumberParser.
type BaseNumberListener struct{}

var _ NumberListener = &BaseNumberListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNumberListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNumberListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNumberListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNumberListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseNumberListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseNumberListener) ExitNumber(ctx *NumberContext) {}
