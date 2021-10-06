# listPrimeFactorsInGo
List Prime Factors in Golang

More taking my daughter's homework far too seriously

Find the prime factors of a number and represent them in 'index format'

If invoked without a parameter it assumes that you want to look at the prime factors of 100.

./listPrimeFactors <positive_number>

It lists values found in index format.


    ./listPrimeFactors 36044456

    2³ × 7 × 643651

## To build

go build -o listPrimeFactors cmd/listPrimeFactors/main.go