//
//
// Normal Distribution sample tests
//
// Author:   Gary Boone
// 
// Copyright (c) 2011 Gary Boone <gary.boone@gmail.com>.
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

//
//
// Benchmark tests
//
// run with: gotest -bench="Benchmark"
//

func BenchmarkBoxMullerTransformation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNormalBMT()
	}
}


func BenchmarkPolarBoxMuller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNormalPolarBM()
	}
}


func BenchmarkZiggurat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNormalZig()
	}
}