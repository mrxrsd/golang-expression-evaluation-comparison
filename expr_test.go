package main

import (
	"testing"

	"github.com/antonmedv/expr"
)

func Benchmark_expr(b *testing.B) {
	params := createParams()

	program, err := expr.Compile(example, expr.Env(params))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func BenchmarkEvaluationNumericLiteral_expr(bench *testing.B) {

	expression, _ := expr.Compile("(2) > (1)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, nil)
	}
}

func BenchmarkEvaluationLiteralModifiers_expr(bench *testing.B) {

	expression, _ := expr.Compile("(2) + (2) == (4)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, nil)
	}
}

func BenchmarkEvaluationParameter_expr(bench *testing.B) {

	parameters := map[string]interface{}{
		"requests_made": 99.0,
	}

	expression, _ := expr.Compile("requests_made", expr.Env(parameters))

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, parameters)
	}
}

func BenchmarkEvaluationParameters_expr(bench *testing.B) {

	parameters := map[string]interface{}{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	expression, _ := expr.Compile("requests_made > requests_succeeded", expr.Env(parameters))

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, parameters)
	}
}

func BenchmarkEvaluationParametersModifiers_expr(bench *testing.B) {

	parameters := map[string]interface{}{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	expression, _ := expr.Compile("(requests_made * requests_succeeded / 100) >= 90", expr.Env(parameters))

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, parameters)
	}
}

func BenchmarkComplexPrecedenceMath_expr(bench *testing.B) {

	expression, _ := expr.Compile("1+2-3*4/5+6-7*8/9+0")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, nil)
	}
}

func BenchmarkMath_expr(bench *testing.B) {

	params := map[string]interface{}{
		"var1": 2,
		"age":  4,
	}

	expression, _ := expr.Compile("var1+2*(3*age)", expr.Env(params))

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expr.Run(expression, params)
	}
}
