package main

const example = `(Origin == 1 || Country == 55) && (Value >= 100 || Adults == 1)`

func createParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["Origin"] = 1
	params["Country"] = 51
	params["Adults"] = 1
	params["Value"] = 100
	return params
}

type Params struct {
	Origin  int
	Country int
	Value   int
	Adults  int
}
