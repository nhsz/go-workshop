package main

import (
	"./divisible"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	invalidNumberOfArguments, argError := hasInvalidNumberOfArguments()

	if !invalidNumberOfArguments {
		if p, err := strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(divisible.SumDivisibleValuesInRange(p, 3, 5))
		}
	} else {
		fmt.Println(argError)
	}
}

func hasInvalidNumberOfArguments() (bool, error) {
	if len(os.Args) != 2 {
		return true, errors.New("Error: you must provide exactly 1 integer argument")
	}

	return false, nil
}
