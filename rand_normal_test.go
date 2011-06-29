//
//
// Normal Distribution function
//
//
package stats

import (
	"testing"
)

func TestBoxMullerTransformation(t *testing.T) {
	var y1, y2 = BoxMullerTransformation(0.3, 0.2)
	checkFloat64(y1, 0.47951886809696076, TOL, "BoxMullerTransformation y1", t)
	checkFloat64(y2, 1.475807326106928, TOL, "BoxMullerTransformation y2", t)
}

