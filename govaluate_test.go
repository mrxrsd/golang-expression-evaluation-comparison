package main

import (
	"testing"

	"github.com/Knetic/govaluate"
)

func Benchmark_govaluate(b *testing.B) {
	params := createParams()

	expression, err := govaluate.NewEvaluableExpression(example)

	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expression.Evaluate(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func BenchmarkEvaluationNumericLiteral_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("(2) > (1)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(nil)
	}
}

func BenchmarkEvaluationLiteralModifiers_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("(2) + (2) == (4)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(nil)
	}
}

func BenchmarkEvaluationParameter_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("requests_made")

	parameters := map[string]interface{}{
		"requests_made": 99.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(parameters)
	}
}

func BenchmarkEvaluationParameters_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("requests_made > requests_succeeded")

	parameters := map[string]interface{}{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(parameters)
	}
}

func BenchmarkEvaluationParametersModifiers_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")

	parameters := map[string]interface{}{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(parameters)
	}
}

func BenchmarkComplexPrecedenceMath_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("1+2-3*4/5+6-7*8/9+0")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(nil)
	}
}

func BenchmarkMath_govaluate(bench *testing.B) {

	expression, _ := govaluate.NewEvaluableExpression("var1+2*(3*age)")

	params := map[string]interface{}{
		"var1": 2,
		"age":  4,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression.Evaluate(params)
	}
}
