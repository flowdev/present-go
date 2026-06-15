package main

import (
	"fmt"

	"github.com/flowdev/comb"
	"github.com/flowdev/comb/cmb"
)

func main() {
	// 1 OMIT
	exprp := cmb.Expression(comb.SafeSpot(cmb.Int64(false, 10))).AddPrefixLevel(cmb.PrefixOp[int64]{
		Op:       "-",
		SafeSpot: false,
		Fn: func(i int64) int64 { return -i },
	}).AddInfixLevel(cmb.InfixOp[int64]{
		Op:       "*",
		SafeSpot: true,
		Fn: func(a, b int64) int64 { return a * b },
	}, cmb.InfixOp[int64]{
		Op:       "/",
		SafeSpot: true,
		Fn: func(a, b int64) int64 { return a / b },
	}).AddInfixLevel(cmb.InfixOp[int64]{
		Op:       "+",
		SafeSpot: true,
		Fn: func(a, b int64) int64 { return a + b },
	}, cmb.InfixOp[int64]{
		Op:       "-",
		SafeSpot: false,
		Fn: func(a, b int64) int64 { return a - b },
	}).AddParentheses("(", ")", true).AddParentheses("[", "]", true).Parser()

	eqp := cmb.Sequence(cmb.Whitespace0(), comb.SafeSpot(cmb.String("=")), cmb.Whitespace0())
	parser := cmb.Map4(comb.SafeSpot(cmb.Alpha1()), cmb.Alphanumeric0(), eqp, exprp,
		func(idStart, idEnd string, eq []string, num int64) (bool, error) {
			fmt.Printf("Statement: %s = %d\n", idStart+idEnd, num)
			return true, nil
		})
	result, err := comb.RunOnString("abc = -2 * 3 + 8 / 1 - 1", parser)
	fmt.Printf("Result: %t, err = %v\n", result, err)
	// 2 OMIT
}
