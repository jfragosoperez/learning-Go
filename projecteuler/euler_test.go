package projecteuler

import(
	"testing"
	"reflect"
)

func TestFindMultiples(t *testing.T) {
	const inMultipleOf, inUpTo, upperBoundForcingEmptyList = 5, 52, 1
	outMultiples := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	result := FindMultiples(inMultipleOf, inUpTo)
	//non multiple-list
	if equal := reflect.DeepEqual(outMultiples, result); !equal {
		t.Errorf("Multiple list of %v upTo %v is %v", inMultipleOf, inUpTo, outMultiples)
	}
	//empty multiple list
	if len(FindMultiples(inMultipleOf, upperBoundForcingEmptyList)) != 0 {
		t.Errorf("Multiple list of %v upTo %v should be empty ([])", inMultipleOf, upperBoundForcingEmptyList)
	}
}