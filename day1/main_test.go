package main

import (
	"./divisible"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := divisible.SumDivisibleValuesInRange(5, 3, 5)
	if sum != 8 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 8)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		divisible.SumDivisibleValuesInRange(100, 3, 5)
	}
}

func ExampleSum() {
	sum := divisible.SumDivisibleValuesInRange(5, 3, 5)
	fmt.Println(sum)
	// Output: 8
}
