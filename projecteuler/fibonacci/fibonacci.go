//fibonacci sequence methods
package fibonacci

import "math"

//Fibonacci calculates the value for the nth term of the fib. sequence.
func Fibonacci(nthTerm uint32)(nthThermValue uint32) {
	return fibonacciRecursive(nthTerm)
}

//Fibonacci recursive
func fibonacciRecursive(nthTerm uint32)(nthThermValue uint32) {
	if nthTerm == 0 {
		return 0
	} else if nthTerm == 1 {
		return 1
	}
	return fibonacciRecursive(nthTerm - 1) + fibonacciRecursive(nthTerm - 2)
}

//Binet's formula appeared from the question: Can we find a formula for F(n) which involves 
//only n and does not need any other (earlier) Fibonacci values?
//we have to math.Floor the value summing 0.5 before casting to int, because 
//GoLang truncates (discards the fraction) instead of rounding
//https://code.google.com/p/go/issues/detail?id=3804
//https://groups.google.com/forum/#!topic/golang-nuts/n2DlRmMlGJ0
func fibonacciBinet(nthTerm uint32)(nthTermValue uint32) {
	squareRootOfFive := math.Sqrt(5)
	phi := (1 + squareRootOfFive) / 2
	return uint32(math.Floor(
		((math.Pow(phi,float64(nthTerm)) - math.Pow((1-phi),float64(nthTerm)))/squareRootOfFive)+0.5)) 
}