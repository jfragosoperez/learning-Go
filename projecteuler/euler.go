//exercises from https://projecteuler.net/problems
package projecteuler

import "errors"

//Calculates the sum of all the numbers we get
//from finding the multiples given a multiplesOf list and
//the limit upTo which is the bound to search the multiples
//from each multiple number in the list.
func SumMultiples(multiplesOf []int, upTo int)(sum int, err error) {
	multiplesSum := 0
	if multiplesOf == nil {
		return -1, errors.New("The multiplesOf list is nil")
	}
	if emptyMultList := len(multiplesOf) == 0; !emptyMultList {
		
	}
	return multiplesSum, nil
}

//Finds all the numbers that are multiple of multipleOf
//given the upTo number as the upper bound.
func FindMultiples(multipleOf int, upTo int) ([]int) {
	numMultiples := upTo / multipleOf
	multiples := make([]int, numMultiples)
	for i := 1; i <= numMultiples; i++ {
		multiples[i-1] = multipleOf * i
	}
	return multiples
}