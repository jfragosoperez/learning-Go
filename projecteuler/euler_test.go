package projecteuler

import(
	"fmt"
	"testing"
)

func TestFindMultiples(t *testing.T) {
	const inMultipleOf, inUpTo, upperBoundForcingEmptyList = 5, 50, 1
	outMultiples := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	result := FindMultiples(inMultipleOf, inUpTo)
	//non multiple-list
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
		if result[i] != outMultiples[i] {
			t.Errorf("Multiple list of %v upTo %v is %v", inMultipleOf, inUpTo, outMultiples)
			break
		}
	}
	//empty multiple list
	if len(FindMultiples(inMultipleOf, upperBoundForcingEmptyList)) != 0 {
		t.Errorf("Multiple list of %v upTo %v should be empty ([])", inMultipleOf, upperBoundForcingEmptyList)
	}
}