//fibonacci sequence methods
package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {

	inNthTerms := [11]uint32{0,1,2,3,20,22,24,26,30,35,40}
	outNthTermsValues := [11]uint32{0,1,1,2,6765,17711,46368,121393,832040,9227465,102334155}

	for i := 0; i < len(inNthTerms); i++ {
		testNthTermShouldBe(inNthTerms[i], outNthTermsValues[i], t)
	}		
}

func testNthTermShouldBe(nthTerm, expectedValue uint32, t *testing.T) {
	if result := Fibonacci(nthTerm); result != expectedValue {
		t.Errorf("The value of the %vth term should be %v," +
			"but the result is %v", nthTerm, expectedValue, result)
	}
}