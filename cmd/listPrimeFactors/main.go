package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/matthewvcarey1/listPrimeFactorsInGo/internal/pkg/sieve"
)

type baseExponent struct {
	base     int
	exponent int
}

func main() {
	t1 := time.Now()
	defer func() {
		taken := time.Now().Sub(t1)
		fmt.Println("Time taken", taken)
	}()
	num := 100
	if len(os.Args) > 1 {
		numStr := os.Args[1]
		ns, err := strconv.Atoi(numStr)
		if err == nil && ns > 0 {
			num = ns
		}
	}
	primes := sieve.SieveOfEratosthenes(num / 2)
	var factors []int
	factors = listPrimeFactors(num, primes, factors)
	sort.Ints(factors)
	nps := makeListOfBaseExponents(factors)

	var outStr: string
	for i, np in range nps {
		if i > 0 {
			outStr = outStr + " \u00D7 "
		}
		outStr = outStr + np.base
		if(np.exponent > 1){
			outStr = outStr + numToSuperscript(np.exponent)
		}
	}
	fmt.Printf(outStr)

}

func listPrimeFactors(num int, primes []int, res []int) []int {
	if num < 2 {
		return res
	}
	for _, p := range primes {
		if num%p == 0 {
			res = append(res, p)
			return listPrimeFactors(num/p, primes, res)
		}
	}
	return res
}

func makeListOfBaseExponents(nums []int) []baseExponent {
	var np []baseExponent
	lastNum := -1
	npIndex := -1
	for _, n := range nums {
		if len(nums) > 0 && n == lastNum {
			np[npIndex].exponent++
		} else {
			npIndex++
			np = append(np, baseExponent{base: n, exponent: 1})
			lastNum = n
		}
	}
	return np
}

func numToSuperscript(num int) string {
	superscriptDigits := [10]string{
		"\u2070",
		"\u00B9",
		"\u00B2",
		"\u00B3",
		"\u2074",
		"\u2075",
		"\u2076",
		"\u2077",
		"\u2078",
		"\u2079",
	}
	var supers string
	for num > 0 {
		val := num % 10
		s := superscriptDigits[val]
		supers = s + supers
	}
	return supers
}
