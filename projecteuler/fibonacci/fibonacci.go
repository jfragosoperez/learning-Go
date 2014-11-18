//fibonacci sequence methods
package fibonacci

import "math"

//Fibonacci calculates the value for the nth term of the fib. sequence.
func Fibonacci(nthTerm uint32)(nthThermValue uint32) {
	return FibonacciBinet(nthTerm)
}

//Returns the Fibonacci Sequence values up to theg nthTerm.
func FibonacciSequence(nthTerm uint32)(sequenceTillNthTherm []uint32){
	sequence := []uint32{}
	
	if nthTerm == 0 { //if nth = 0 -> return []
		return sequence
	} 
	//add first element
	sequence = append(sequence, 1)

	if(nthTerm == 1 || nthTerm == 2){ //return [1] or [1,1] in case of nthTerm is 1 or two
		if nthTerm == 2 { // add one more el, in case of 2nd therm
			sequence = append(sequence, 1)
		} 
		return sequence
	}
	sequence = append(sequence, 1)
	
	for ith := 3;  ith <= int(nthTerm); ith++ { //case nth > 2
		//seq[(ith-1)-1] + seq[(ith-2)-1]
		sequence = append(sequence, sequence[ith-2] + sequence[ith-3])
	}
	return sequence
}

//Fibonacci iterative
func FibonacciIterative(nthTerm uint32)(nthTermValue uint32) {
	var tmp uint32
	previousValues := make([]uint32, 2)

	if nthTerm == 0 { //if nth = 0 
		return 0
	} 

	if(nthTerm == 1 || nthTerm == 2){ //nthTerm is 1 or two
		return 1
	}

	previousValues[0] = 1
	previousValues[1] = 1

	for ith := 3;  ith <= int(nthTerm); ith++ { //case nth > 2
		tmp = previousValues[0] + previousValues[1]
		previousValues[0] = previousValues[1]
		previousValues[1] = tmp
	}
	return previousValues[1]
}

//Fibonacci recursive
func FibonacciRecursive(nthTerm uint32)(nthThermValue uint32) {
	if nthTerm == 0 {
		return 0
	} else if nthTerm == 1 {
		return 1
	}
	return FibonacciRecursive(nthTerm - 1) + FibonacciRecursive(nthTerm - 2)
}

//Binet's formula appeared from the question: Can we find a formula for F(n) which involves 
//only n and does not need any other (earlier) Fibonacci values?
//we have to math.Floor the value summing 0.5 before casting to int, because 
//GoLang truncates (discards the fraction) instead of rounding. Do not forget
//that this is an approximation (because phi and squareRootOfFive here are finite).
//https://code.google.com/p/go/issues/detail?id=3804
//https://groups.google.com/forum/#!topic/golang-nuts/n2DlRmMlGJ0
func FibonacciBinet(nthTerm uint32)(nthTermValue uint32) {
	squareRootOfFive := math.Sqrt(5)
	phi := (1 + squareRootOfFive) / 2
	return uint32(math.Floor(
		((math.Pow(phi,float64(nthTerm)) - math.Pow((1-phi),float64(nthTerm)))/squareRootOfFive)+0.5)) 
}