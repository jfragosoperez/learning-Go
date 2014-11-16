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

func TestSumMultiples(t *testing.T){
	const outSumEmptyList = 0
	const inUpTo = 100
	inEmptyList := []int{}
	if result, error := SumMultiples(inEmptyList, inUpTo); error==nil && result != 0 {
		t.Errorf("When giving an empty multiple list, the expected sum should be 0")
	}
	//in order to avoid compilation error (declared and not used error) 
	//we use the blank identifier _ : It serves as an anonymous placeholder 
	//instead of a regular (non-blank) identifier and has special meaning 
	//in declarations, as an operand, and in assignments.
	_, error := SumMultiples(nil, inUpTo)
	if error == nil {
		t.Errorf("When given a nil multiplesOf list, error should be different than nil")
	}
}