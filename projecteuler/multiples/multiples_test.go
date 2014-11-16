package multiples

import(
	"testing"
	"reflect"
)

func TestFindMultiples(t *testing.T) {
	const inMultipleOf, inUpTo, upperBoundForcingEmptyList = 5, 50, 1
	outMultiples := []int{5, 10, 15, 20, 25, 30, 35, 40, 45}
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
	const inUpTo100 = 100
	inEmptyList := []int{}
	if result, error := SumMultiples(inEmptyList, inUpTo100); error == nil && result != 0 {
		t.Errorf("When giving an empty multiple list, the expected sum should be 0")
	}
	//in order to avoid compilation error (declared and not used error) 
	//we use the blank identifier _ : It serves as an anonymous placeholder 
	//instead of a regular (non-blank) identifier and has special meaning 
	//in declarations, as an operand, and in assignments.
	_, error := SumMultiples(nil, inUpTo100)
	if error == nil {
		t.Errorf("When given a nil multiplesOf list, error should be different than nil")
	}

	inMultiplesOf2and3 := []int{2, 3}
	const inUpTo10 = 10
	//(2+4+6+8+3+9) = 32
	const expectedSum = 32
	if result, error := SumMultiples(inMultiplesOf2and3, inUpTo10); error != nil || result != expectedSum {
		t.Errorf("Multiples from %v upTo %v would give no error and as result: %v." +
			" But our result has been error = %v and result = %v", 
			inMultiplesOf2and3, inUpTo10, expectedSum, error, result)
	}
}