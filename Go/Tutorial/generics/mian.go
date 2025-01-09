package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64 {
		"first": 32,
		"second": 12,
	}

	floats := map[string]float64 {
		"first":  35.98,
        "second": 26.99,
	}

	fmt.Printf(
		"Non-Generics Sums: %v, %v\n", 
		SumInts(ints), 
		SumFloats(floats))

	fmt.Printf(
		"Generics Sums: %v, %v\n",
		SumIntsorOrFloats[string, int64](ints),
		SumIntsorOrFloats[string, float64](floats))
	
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
	SumIntsorOrFloats(ints),
	SumIntsorOrFloats(floats))


	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
}

func SumInts(m map[string]int64) int64 {  // m is a map with string keys and int values
	var s int64

	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}
	return s
}


func SumIntsorOrFloats [K comparable, V int64 | float64](m map[K]V) V {  // K is a comparable type, such as string, V is either int64 or float64
	var s V

	for _, v := range m {
		s += v
	}
	return s
}


func SumNumbers [K comparable, V Number] (m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}