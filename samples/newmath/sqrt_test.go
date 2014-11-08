package newmath

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

//to execute tests from a package call -> go test YOUR_PROJECT_LOCATION/
//or go test if you are in the source directory