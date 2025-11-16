package main

import (
	"fmt"

	"github.com/flowdev/comb"
	"github.com/flowdev/comb/cmb"
)

func main() {
	type data struct {
		conf string
		year string
	}

	parser := cmb.Map2(cmb.Alpha1(), cmb.Digit1(), func(s1, s2 string) (data, error) {
		return data{conf: s1, year: s2}, nil
	})
	result, err := comb.RunOnString("DevFest2025", parser)
	fmt.Printf("Conference: %s, Year: %s, err: %v\n", result.conf, result.year, err)
}
