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

func BenchmarkEvaluationNumericLiteral_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})

	formula, _ := engine.Build("(2) > (1)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(nil)
	}
}

func BenchmarkEvaluationLiteralModifiers_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})

	formula, _ := engine.Build("(2) + (2) == (4)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(nil)
	}
}

func BenchmarkEvaluationParameter_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})
	formula, _ := engine.Build("requests_made")
	parameters := map[string]float64{
		"requests_made": 99.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}

func BenchmarkEvaluationParameters_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})
	formula, _ := engine.Build("requests_made > requests_succeeded")
	parameters := map[string]float64{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}

func BenchmarkEvaluationParametersModifiers_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})

	formula, _ := engine.Build("(requests_made * requests_succeeded / 100) >= 90")
	parameters := map[string]float64{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}

func BenchmarkComplexPrecedenceMath_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})

	formula, _ := engine.Build("1+2-3*4/5+6-7*8/9+0")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(nil)
	}
}

func BenchmarkMath_gojacego(bench *testing.B) {

	engine := gojacego.NewCalculationEngineWithOptions(gojacego.JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})

	formula, _ := engine.Build("var1+2*(3*age)")

	params := map[string]float64{
		"var1": 2,
		"age":  4,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(params)
	}
}
