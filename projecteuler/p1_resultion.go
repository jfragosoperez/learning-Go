package main

import(
	"fmt"
	"com.jfragosoperez/projecteuler/multiples"
)

//If we list all the natural numbers below 10 that are 
//multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
func main(){
	const upTo = 1000
	multipleList := []int{3, 5}
	if sumMultiples, error := multiples.SumMultiples(multipleList, upTo); error == nil {
		fmt.Println("The sum of all multiples of 3 or 5 below 1000 is", sumMultiples)
		//result is 233168
	}
}