package main

import (
	"fmt"
	"github.com/dharmamike/syndiv/syndiv"
	"log"
	"os"
	"strconv"
)

func main() {
	cArg := os.Args[1]
	c, err := strconv.Atoi(cArg)
	if err != nil {
		log.Fatal("The value at which the quotient and remainder are calculated must be an integer.")
	}

	coeffArgs := os.Args[2:]
	coeffs := make([]int, len(coeffArgs))

	for i, coeffArg := range coeffArgs {
		coeff, err := strconv.Atoi(coeffArg)
		if err != nil {
			log.Fatal("The coefficients of the polynomial must be integers in descending order.")
		}
		coeffs[i] = coeff
	}

	result := syndiv.Syndiv(coeffs, c)
	fmt.Printf("The coefficients of the quotient are %v and the remainder is %d\n", result[:len(result)-1], result[len(result)-1])
}
