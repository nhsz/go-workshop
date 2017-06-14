package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	p, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sumDivisibleValues(p))
	}
}

func sumDivisibleValues(p int) int {
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
