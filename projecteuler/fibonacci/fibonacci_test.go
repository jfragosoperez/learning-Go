//fibonacci sequence methods
package fibonacci

import(
	"testing"
	"reflect"
)

func TestFibonacci(t *testing.T) {

	inNthTerms := [11]uint32{0,1,2,3,20,22,24,26,30,35,40}
	outNthTermsValues := [11]uint32{0,1,1,2,6765,17711,46368,121393,832040,9227465,102334155}

	for i := 0; i < len(inNthTerms); i++ {
		testNthTermShouldBe(inNthTerms[i], outNthTermsValues[i], t)
	}		
}

func TestFibonacciSequence(t *testing.T){
	outNthTermsValues := []uint32{1,1,2,3,5,8,13,21,34,55,89,144,
		233,377,610,987,1597,2584,4181,6765}

	if len(FibonacciSequence(0)) != 0 {
		t.Errorf("Fibonacci sequence of the 0th term should return an empty list")
	}

	for i := 0; i < len(outNthTermsValues); i++ {
		testNthTermSequence(uint32(i+1), outNthTermsValues[0:(i+1)], t)
	}

}

func testNthTermSequence(nthTerm uint32, expectedList []uint32, t * testing.T) {
	if resultList := FibonacciSequence(nthTerm); !reflect.DeepEqual(resultList, expectedList) {
		t.Errorf("Fibonacci sequence of %vth term would be %v, but is returning %v instead", 
			nthTerm, expectedList, resultList)
	}
}

func testNthTermShouldBe(nthTerm, expectedValue uint32, t *testing.T) {
	if result := Fibonacci(nthTerm); result != expectedValue {
		t.Errorf("The value of the %vth term should be %v," +
			"but the result is %v", nthTerm, expectedValue, result)
	}
}

// TO RUN BENCHMARKS -> go test -bench . -> executes all benchmarks

func BenchmarkFiboIterative(b *testing.B) {
    for i := 0; i < b.N; i++ {
    	FibonacciIterative(uint32(i))
    }
}

func BenchmarkFiboBinet(b *testing.B) {
    for i := 0; i < b.N; i++ {
    	FibonacciBinet(uint32(i))
    }
}

func BenchmarkFiboSequence(b *testing.B) {
    for i := 0; i < b.N; i++ {
    	FibonacciSequence(uint32(i))
    }
}

func BenchmarkFiboRecursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
    	FibonacciRecursive(uint32(i))
    }
}



