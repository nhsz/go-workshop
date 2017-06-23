package divisible

/*
SumDivisibleValuesInRange sums all the v values within
the range 0 <= v <= p, such that v is divisible by x or y
*/
func SumDivisibleValuesInRange(p, x, y int) int {
	sum := 0
	for v := 0; v <= p; v++ {
		if IsDivisibleBy(v, x, y) {
			sum += v
		}
	}

	return sum
}

// IsDivisibleBy checks if the value v is divisible by i and j
func IsDivisibleBy(v, i, j int) bool {
	return v%i == 0 || v%j == 0
}
