package divisible

func SumDivisibleValuesInRange(p int) int {
	sum := 0
	for v := 0; v <= p; v++ {
		if isDivisibleBy(v, 3, 5) {
			sum += v
		}
	}

	return sum
}

func isDivisibleBy(v, i, j int) bool {
	return v%i == 0 || v%j == 0
}
