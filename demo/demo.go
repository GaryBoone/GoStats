//
// demo.go
//
// Author:    Gary Boone
//

package main

import (
	"stats" // assumes you've done 'make install' in the stats directory
	"fmt"
	"rand"
	"time"
)

const NUM_SAMPLES = 5

func incrementalDemo() {
	var d stats.Desc
	fmt.Printf("\n**** Descriptive Statistics, Incremental Updates **\n")
	fmt.Printf("** create a new descriptive stats struct:\n")
	fmt.Printf("var d stats.Desc\n")
	fmt.Printf("** update it with new values:\n")
	for i := 0; i < NUM_SAMPLES; i++ {
		x := rand.Float64()*100.0 - 25.0 // uniform samples in {-25, 75}
		fmt.Printf("d.Update(%v)\n", x)
		d.Update(x)
	}
	fmt.Printf("then request stats, for example:    d.Count()\n")
	fmt.Printf("\n** Descriptive Statistics:\n")
	fmt.Printf("count = %v\n", d.Count())
	fmt.Printf("min = %v\n", d.Min())
	fmt.Printf("max = %v\n", d.Max())
	fmt.Printf("sum = %v\n", d.Sum())
	fmt.Printf("mean = %v\n", d.Mean())
	fmt.Printf("standard deviation = %v\n", d.SampleStandardDeviation())
	fmt.Printf("variance = %v\n", d.SampleVariance())
	skew := d.SampleSkew()
	fmt.Printf("skew = %v\n", skew)
	// 
	// The rand functions return uniformly distributed values. Therefore, with
	// enough samples, the skew should be 0. Try increasing NUM_SAMPLES
	if skew < 0.0 {
		fmt.Printf("  The sample is skews left (longer tail to left).\n")
	} else if skew > 0.0 {
		fmt.Printf("  The sample is skews right (longer tail to right).\n")
	}
	kurtosis := d.SampleKurtosis()
	fmt.Printf("kurtosis = %v\n", kurtosis)
	// 
	// The rand functions return uniformly distributed values. Therefore, with
	// enough of them, our sample should be flat, or platykurtic.
	if kurtosis < 0.0 {
		fmt.Printf("  The sample is platykurtic.\n")
	} else if kurtosis > 0.0 {
		fmt.Printf("  The sample is leptokurtic.\n")
	}
}

func batchDemo() {
	fmt.Printf("\n**** Descriptive Statistics, Batch Updates **\n")
	fmt.Printf("** array of values:\n")
	a := makeArray(NUM_SAMPLES, 100, -25)
	printArray("x", a)
	fmt.Printf("}\n")
	fmt.Printf("then request stats, for example:    stats.StatsCount(a)\n")
	fmt.Printf("** Descriptive Statistics:\n")
	fmt.Printf("count = %v\n", stats.StatsCount(a))
	fmt.Printf("min = %v\n", stats.StatsMin(a))
	fmt.Printf("max = %v\n", stats.StatsMax(a))
	fmt.Printf("sum = %v\n", stats.StatsSum(a))
	fmt.Printf("mean = %v\n", stats.StatsMean(a))
	fmt.Printf("standard deviation = %v\n", stats.StatsSampleStandardDeviation(a))
	fmt.Printf("variance = %v\n", stats.StatsSampleVariance(a))
	fmt.Printf("skew = %v\n", stats.StatsSampleSkew(a))
	fmt.Printf("kurtosis = %v\n", stats.StatsSampleKurtosis(a))
}

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
		xData[i] = float64(i)*3.0
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

	rand.Seed(time.Nanoseconds())

	incrementalDemo()
	batchDemo()
	incrementalRegressionDemo()
	batchRegressionDemo()
}
