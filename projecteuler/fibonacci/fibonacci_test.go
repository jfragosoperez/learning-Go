//fibonacci sequence methods
package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	testNthTermShouldBe(0, 0, t)
	testNthTermShouldBe(1, 1, t)
	testNthTermShouldBe(2, 1, t)
	testNthTermShouldBe(3, 2, t)
	testNthTermShouldBe(20, 6765, t)
}

func testNthTermShouldBe(nthTerm, expectedValue uint32, t *testing.T) {
	if result := Fibonacci(nthTerm); result != expectedValue {
		t.Errorf("The value of the %vth term should be %v," +
			"but the result is %v", nthTerm, expectedValue, result)
	}
}