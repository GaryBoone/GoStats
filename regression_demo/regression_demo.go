//
// regression_demo.go
//
// Can be compiled and run from the main GoStats directory using:
// $ (cd demos/regression/ && make clean && make)
// $ demos/regression/regression_demo
//
// Author:    Gary Boone
//

package main

import (
	"GoStats/stats"
	"fmt"
	"math/rand"
	"time"
)

const NUM_SAMPLES = 5

func incrementalRegressionDemo() {
	var r stats.Regression
	fmt.Printf("\n**** Linear Regression, Incremental Updates **\n")
	fmt.Printf("** create a new regresson struct:\n")
	fmt.Printf("var r stats.Regression\n")
	fmt.Printf("** update it with new values:\n")
	for i := 0; i < NUM_SAMPLES; i++ {
		x := float64(i) * 3.0
		y := rand.Float64()*100.0 - 25.0 // uniform samples in {-25, 75}
		fmt.Printf("r.Update(%v, %v)\n", x, y)
		r.Update(x, y)
	}
	fmt.Printf("then request regression results, for example:    r.Slope()\n")
	fmt.Printf("** Linear Regression:\n")
	fmt.Printf("count = %v\n", r.Count())
	fmt.Printf("slope = %v\n", r.Slope())
	fmt.Printf("intercept = %v\n", r.Intercept())
	fmt.Printf("r-squared = %v\n", r.RSquared())
	fmt.Printf("slope standard error = %v\n", r.SlopeStandardError())
	fmt.Printf("intercept standard error = %v\n", r.InterceptStandardError())
}

func makeArray(size int, width, min float64) []float64 {
	a := make([]float64, size)
	for i := 0; i < size; i++ {
		x := rand.Float64()*width + min // uniform samples in {-25, 75}
		a[i] = x
	}
	return a
}

func printArray(name string, a []float64) {
	fmt.Printf(" %v={", name)
	for i := 0; i < len(a); i++ {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%v", a[i])
	}
	fmt.Printf("}\n")
}

func batchRegressionDemo() {
	fmt.Printf("\n**** Linear Regression, Batch Updates **\n")
	fmt.Printf("** array of values:\n")
	xData := make([]float64, NUM_SAMPLES)
	for i := 0; i < NUM_SAMPLES; i++ {
		xData[i] = float64(i) * 3.0
	}
	printArray("xData", xData)
	yData := makeArray(NUM_SAMPLES, 100, -25)
	printArray("yData", yData)
	var slope, intercept, rsquared, count, slopeStdErr, intcptStdErr = stats.LinearRegression(xData, yData)
	fmt.Printf("then request regression results, for example:\n")
	fmt.Printf("      var slope, intercept, _, _, _, _ = stats.LinearRegression(xData, yData)\n")
	fmt.Printf("** Linear Regression:\n")
	fmt.Printf("slope = %v\n", slope)
	fmt.Printf("intercept = %v\n", intercept)
	fmt.Printf("r-squared = %v\n", rsquared)
	fmt.Printf("count = %v\n", count)
	fmt.Printf("slope standard error = %v\n", slopeStdErr)
	fmt.Printf("intercept standard error = %v\n", intcptStdErr)
}

func main() {
	fmt.Printf("GoStats Demo\n\n")

	rand.Seed(int64(time.Now().Nanosecond()))

	incrementalRegressionDemo()
	batchRegressionDemo()
}
