//fibonacci sequence methods
package fibonacci

import "math"

//Fibonacci calculates the value for the nth term of the fib. sequence.
func Fibonacci(nthTerm uint32)(nthThermValue uint32) {
	return fibonacciBinet(nthTerm)
}

//Returns the Fibonacci Sequence till the nthTerm.
func FibonacciSequence(nthTerm uint32)(sequenceTillNthTherm []uint32){
	sequence := []uint32{}
	//if nth = 0 -> return []
	if nthTerm == 0 {
		return sequence
	} 
	//add first element
	sequence = append(sequence, 1)
	//return [1] or [1,1] in case of nthTerm is 1 or two
	if(nthTerm == 1 || nthTerm == 2){
		if nthTerm == 2 { 
			sequence = append(sequence, 1)
		} 
		return sequence
	}
	sequence = append(sequence, 1)
	//case nth > 2
	for ith := 3;  ith <= int(nthTerm); ith++ {
		//seq[(ith-1)-1] + seq[(ith-2)-1]
		sequence = append(sequence, sequence[ith-2] + sequence[ith-3])
	}
	return sequence
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