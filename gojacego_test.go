package main

import (
	"testing"

	"github.com/mrxrsd/gojacego"
)

func Benchmark_gojacego(b *testing.B) {
	p := createParams()

	params := map[string]float64{
		"Origin":  float64(p["Origin"].(int)),
		"Country": float64(p["Country"].(int)),
		"Value":   float64(p["Value"].(int)),
		"Adults":  float64(p["Adults"].(int)),
	}

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})

	expression, err := engine.Build(example)

	if err != nil {
		b.Fatal(err)
	}

	var out float64

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out = expression(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out != 1.0 {
		b.Fail()
	}
}
