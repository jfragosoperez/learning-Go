//exercises from https://projecteuler.net/problems
package multiples

import "errors"

//Calculates the sum of all the numbers we get
//from finding the multiples given a multiplesOf list and
//the limit upTo which is the upper bound (not included). Thus the multiple
//search will find values in the set [multiplesOf[i], upTo)
//We want to guarantee mutual exclusion, so the sum does not include
//repeated multiples.
func SumMultiples(multiplesOf []int, upTo int)(sum int, err error) {
	multiplesSum := 0
	if multiplesOf == nil {
		return -1, errors.New("The multiplesOf list is nil")
	}
	
	if multiplesOfLength := len(multiplesOf); multiplesOfLength != 0 {
		for i := 0 ; i < upTo ; i++ {
			for z := 0; z < multiplesOfLength; z++ {
				if added := sumIfMultiple(i, multiplesOf[z], &multiplesSum); added {
					break
				}
			}
		}
	}
	
	return multiplesSum, nil
}

func sumIfMultiple(value, multipleOf int, sum *int)(added bool){
	if value % multipleOf == 0 {
		*sum += value
		return true
	}
	return false
}

//Finds all the numbers that are multiple of multipleOf
//given the upTo number as the upper bound (not included).
func FindMultiples(multipleOf int, upTo int) ([]int) {
	numMultiples := (upTo-1) / multipleOf
	multiples := make([]int, numMultiples)
	for i := 1; i <= numMultiples; i++ {
		multiples[i-1] = multipleOf * i
	}
	return multiples
}