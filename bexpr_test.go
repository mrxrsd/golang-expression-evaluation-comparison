package main

import (
	"testing"

	"github.com/hashicorp/go-bexpr"
)

func Benchmark_bexpr(b *testing.B) {
	p := createParams()
	params := Params{
		Origin:  p["Origin"].(int),
		Country: p["Country"].(int),
		Value:   p["Value"].(int),
		Adults:  p["Adults"].(int),
	}

	// Replace operators and parentheses as bexpr can't parse them correctly. So sad :(
	eval, err := bexpr.CreateEvaluator(
		`Origin == 1 and Country == 51 and Value == 100 and Adults == 1`)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = eval.Evaluate(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
