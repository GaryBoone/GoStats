//
// descriptive_stats_demo.go
//
//
// Can be compiled and run from the main GoStats directory using:
// $ (cd demos/descriptive_statistics/ && make clean && make)
// $ demos/descriptive_statistics/descriptive_stats_demo
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

func incrementalDemo() {
	var d stats.Stats
	fmt.Printf("\n**** Descriptive Statistics, Incremental Updates **\n")
	fmt.Printf("** create a new descriptive stats struct:\n")
	fmt.Printf("var d stats.Stats\n")
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

func normalDistributionDemo() {
	maxTrials := 100
	printEvery := 10
	fmt.Printf("\nGenerating %v normal samples. The descriptive statistics ", maxTrials)
	fmt.Printf("converge to the appropriate values: mean→0, variance→1, skew→0, kurtosis→0. \n")
	var d stats.Stats
	for i := 0; i <= maxTrials; i++ {
		y := rand.NormFloat64()
		d.Update(y)
		if i != 0 && i%printEvery == 0 {
			mean := d.Mean()
			variance := d.PopulationVariance()
			skew := d.PopulationSkew()
			kurtosis := d.PopulationKurtosis()
			fmt.Printf("itr %v: mean (→0.0) = %0.5f, variance (→1.0) = %0.5f, skew (→0.0) = %0.5f, kurtosis (→0.0) = %0.5f\n",
				i, mean, variance, skew, kurtosis)
		}
	}
}

func main() {
	fmt.Printf("GoStats Demo\n\n")

	rand.Seed(int64(time.Now().Nanosecond()))

	incrementalDemo()
	batchDemo()
	normalDistributionDemo()
}
