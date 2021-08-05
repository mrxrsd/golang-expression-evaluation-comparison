package main

import (
	"context"
	"testing"

	"github.com/PaesslerAG/gval"
)

func Benchmark_gval(b *testing.B) {
	params := createParams()
	ctx := context.Background()

	var out interface{}
	var err error

	eval, err := gval.Full().NewEvaluable(example)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = eval(ctx, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func BenchmarkEvaluationNumericLiteral_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("(2) > (1)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), nil)
	}
}

func BenchmarkEvaluationLiteralModifiers_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("(2) + (2) == (4)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), nil)
	}
}

func BenchmarkEvaluationParameter_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("requests_made")

	parameters := map[string]interface{}{
		"requests_made": 99.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), parameters)
	}
}

func BenchmarkEvaluationParameters_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("requests_made > requests_succeeded")

	parameters := map[string]interface{}{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), parameters)
	}
}

func BenchmarkEvaluationParametersModifiers_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("(requests_made * requests_succeeded / 100) >= 90")

	parameters := map[string]interface{}{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), parameters)
	}
}

func BenchmarkComplexPrecedenceMath_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("1+2-3*4/5+6-7*8/9+0")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), nil)
	}
}

func BenchmarkMath_gval(bench *testing.B) {

	expression, _ := gval.Full().NewEvaluable("var1+2*(3*age)")

	params := map[string]interface{}{
		"var1": 2,
		"age":  4,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		expression(context.Background(), params)
	}
}
