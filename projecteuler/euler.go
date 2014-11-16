//exercises from https://projecteuler.net/problems
package projecteuler

//Calculates the sum of all the numbers we get
//from finding the multiples given a multiplesOf list and
//the limit upTo which is the bound to search the multiples
//from each multiple number in the list.
func SumMultiples(multiplesOf []int, upTo int) {
	
}

//Finds all the numbers that are multiple of multipleOf
//given the upTo number as the upper bound.
func FindMultiples(multipleOf int, upTo int) ([]int) {
	//case upper bound is less than the multiple
	//the multiple list is empty
	if emptyList := upTo < multipleOf; emptyList {
		return make([]int, 0)
	}
	numMultiples := upTo / multipleOf
	multiples := make([]int, numMultiples)
	for i := 1; i <= numMultiples; i++ {
		multiples[i-1] = multipleOf * i
	}
	return multiples
}