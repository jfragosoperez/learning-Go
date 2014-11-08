//trivial example of a go library
package newmath

//Aprox. to the root square of x (Newton's method)
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
} 

//to compile this package call -> go build YOUR_PROJECT_LOCATION/ or
//go build if you are in the source directory