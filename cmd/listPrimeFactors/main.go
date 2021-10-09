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
		fmt.Println("\n\nTime taken", taken)
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
	lbe := makeListOfBaseExponents(factors)
	outStr := generateIndexFormatString(lbe)

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
	var lbe []baseExponent
	lastNum := -1
	lbeIndex := -1
	for i, n := range nums {
		if i > 0 && n == lastNum {
			lbe[lbeIndex].exponent++
		} else {
			lbeIndex++
			lbe = append(lbe, baseExponent{base: n, exponent: 1})
			lastNum = n
		}
	}
	return lbe
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
		num = num / 10
	}
	return supers
}

func generateIndexFormatString(lbe []baseExponent) string {
	var outStr string
	for i, be := range lbe {
		if i > 0 {
			outStr = outStr + " \u00D7 "
		}
		outStr = outStr + strconv.Itoa(be.base)
		if be.exponent > 1 {
			outStr = outStr + numToSuperscript(be.exponent)
		}
	}
	return outStr
}
