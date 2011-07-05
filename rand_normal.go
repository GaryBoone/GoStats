//
//
// Normal Distribution Random Value functions
//
// Author:   Gary Boone
// 
// Copyright (c) 2011 Gary Boone <gary.boone@gmail.com>.
//
// RandNormal() returns a random value from a normal distribution.
//
// This package contains three methods for generating random values that are normally
// distributed. All are exact methods, differing in implmentation and speed. Run the
// benchmarks to see the speeds: gotest -bench="Benchmark"
//
// The three functions are
//	   RandNormalBMT()             // Box Muller Transformation
//	   RandNormalPolarBM()         // Polar Box Muller
//	   RandNormalZig()             // Ziggurat Method of JA Doornik
//
// Each method generates one random value. The BMT and PBM methods also have methods
// that return a pair of values. Just suffix the function name with 'Pair' 
//
// RandNormal() is an alias to RandNormalPolarBM(), the fastest method by benchmark
//

package stats

import (
	"math"
	"rand"
	"time"
)

var use_spare bool
var spare float64

func init() {
	rand.Seed(time.Nanoseconds())
	zigNorInit(ZIGNOR_C, ZIGNOR_R, ZIGNOR_V)
}

func BoxMullerTransformation(x1, x2 float64) (y1, y2 float64) {
	y1 = math.Sqrt(-2.0*math.Log(x1)) * math.Cos(2.0*math.Pi*x2)
	y2 = math.Sqrt(-2.0*math.Log(x1)) * math.Sin(2.0*math.Pi*x2)
	return
}


// Generate 2 Normal distribution samples from uniform samples using the 
// Box-Muller Transformation.
func RandNormalBMTPair() (float64, float64) {
	x1 := rand.Float64()
	x2 := rand.Float64()
	return BoxMullerTransformation(x1, x2)
}

// Generate a random normal using the Box Muller Transformation, memoizing the second value
func RandNormalBMT() (r float64) {
	if !use_spare {
		r, spare = RandNormalBMTPair()
		use_spare = true
	} else {
		r = spare
		use_spare = false
	}
	return
}

// return a random value v, uniformly drawn from {-1,1}, v != 0
func RandUniformM1To1() (v float64) {
	for ; v == 0.0; v = 2.0*rand.Float64() - 1.0 {
	}
	return
}

// Generate 2 Normal distribution samples from uniform samples using the 
// Polar Box-Muller Transformation, which avoids the Sin and Cos fns.
func RandNormalPolarBMPair() (float64, float64) {
	var x, y, d float64
	for {
		x = RandUniformM1To1()
		y = RandUniformM1To1()
		d = x*x + y*y
		if d < 1.0 {
			break
		}
	}
	f := math.Sqrt(-2.0 * math.Log(d) / d)
	return f * x, f * y
}

// Generate a random normal using the Polar Box Muller, memoizing the second value
func RandNormalPolarBM() (r float64) {
	if !use_spare {
		r, spare = RandNormalPolarBMPair()
		use_spare = true
	} else {
		r = spare
		use_spare = false
	}
	return
}

// ZIGNOR
//
// http://www.doornik.com/research/ziggurat.pdf
//
// An Improved Ziggurat Method to Generate Normal Random Samples
// by JA DOORNIK
//

const ZIGNOR_C = 128            /* number of blocks */
const ZIGNOR_R = 3.442619855899 /* start of the right tail */
/* (R * phi(R) + Pr(X>=R)) * sqrt(2\pi) */
const ZIGNOR_V = 9.91256303526217e-3
/* s_adZigX holds coordinates, such that each rectangle has*/
/* same area; s_adZigR holds s_adZigX[i + 1] / s_adZigX[i] */
var s_adZigX [ZIGNOR_C + 1]float64
var s_adZigR [ZIGNOR_C]float64

func DRanNormalTail(dMin float64, iNegative bool) float64 {
	var x float64
	for {
		x = math.Log(rand.Float64()) / dMin
		y := math.Log(rand.Float64())
		if -2.0*y >= x*x {
			break
		}
	}
	if iNegative {
		return x - dMin
	}
	return dMin - x
}


func zigNorInit(iC int, dR, dV float64) {
	f := math.Exp(-0.5 * dR * dR)
	s_adZigX[0] = dV / f /* [0] is bottom block: V / f(R) */
	s_adZigX[1] = dR
	s_adZigX[iC] = 0
	for i := 2; i < iC; i++ {
		s_adZigX[i] = math.Sqrt(-2.0 * math.Log(dV/s_adZigX[i-1]+f))
		f = math.Exp(-0.5 * s_adZigX[i] * s_adZigX[i])
	}
	for i := 0; i < iC; i++ {
		s_adZigR[i] = s_adZigX[i+1] / s_adZigX[i]
	}
}


func RandNormalZig() (x float64) {
	for {
		u := 2.0*rand.Float64() - 1.0
		i := rand.Uint32() & 0x7F
		/* first try the rectangular boxes */
		if math.Fabs(u) < s_adZigR[i] {
			return u * s_adZigX[i]
		}
		/* bottom box: sample from the tail */
		if i == 0 {
			return DRanNormalTail(ZIGNOR_R, u < 0)
		}
		/* is this a sample from the wedges? */
		x = u * s_adZigX[i]
		f0 := math.Exp(-0.5 * (s_adZigX[i]*s_adZigX[i] - x*x))
		f1 := math.Exp(-0.5 * (s_adZigX[i+1]*s_adZigX[i+1] - x*x))
		if f1+rand.Float64()*(f0-f1) < 1.0 {
			break
		}
	}
	return
}


func RandNormal() (x float64) {
	return RandNormalZig()
}