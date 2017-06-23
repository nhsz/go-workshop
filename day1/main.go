package main

import (
	"./divisible"
	"fmt"
	"os"
	"strconv"
)

func main() {
	p, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(divisible.SumDivisibleValuesInRange(p, 3, 5))
	}
}
