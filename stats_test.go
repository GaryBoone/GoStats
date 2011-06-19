package stats

// 
// stats.go testing
//
// To test, all code was compared against the R stats package (http://r-project.org)
//
// R test code (set 'a' appropriately for each test):
//   a=c(1,2,3,4,5);n=length(a);n;min(a);max(a);sum(a);mean(a);((n-1)/n)*var(a);var(a);
//   sqrt(var(a)*(n-1)/n);sd(a);skewness(a);sqrt(n*(n-1))/(n-2)*skewness(a);
//   kurtosis(a)-3;(n-1)/(n-2)/(n-3)*((n+1)*(kurtosis(a)-3)+6)
//
// The above prints out a list of values corresponding to the list of test results shown 
// in the tests below.
// 
// Unpacking:
//   a=c(1,2,3,4,5);                               // creates a list of input values
//   n=length(a);n;                                // print the Count()
//   min(a); max(a); sum(a); mean(a);	           // Min(), Max(), Sum(), Mean()
//   ((n-1)/n)*var(a);                             // PopulationVariance()
//   var(a);                                       // SampleVariance()
//   sqrt(var(a)*(n-1)/n);                         // PopulationStandardDeviation()
//   sd(a);                                        // SampleStandardDeviation()
//   skewness(a);                                  // PopulationSkew()
//   sqrt(n*(n-1))/(n-2)*skewness(a);              // SampleSkew()
//   kurtosis(a)-3;                                // PopulationKurtosis()
//   (n-1)/(n-2)/(n-3)*((n+1)*(kurtosis(a)-3)+6)   // SampleKurtosis()
//
// R Notes:
//  var() returns the sample variance.
//  sd() returns sample standard deviation.
//  skewness() returns the population skew.
//  kurtosis() returns the population kurtosis, which is not the excess pop kurtosis.
//

import (
	"testing"
	"math"
	"rand"
	"time"
)

const TOL = 1e-14


//
//
// Incremental stats tests
//
//

// With no updates, these are the results on initialization
func TestAppend0(t *testing.T) {
	var d Desc
	checkInt(d.Count(), 0, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkNaN(d.PopulationVariance(), "PopulationVariance", t)
	checkNaN(d.SampleVariance(), "SampleVariance", t)
	checkNaN(d.PopulationStandardDeviation(), "PopulationStandardDeviation", t)
	checkNaN(d.SampleStandardDeviation(), "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 1 value
func TestAppend1(t *testing.T) {
	var d Desc
	d.Append(2.3)
	checkInt(d.Count(), 1, "Count", t)
	checkFloat64(d.Min(), 2.3, TOL, "Min", t)
	checkFloat64(d.Max(), 2.3, TOL, "Max", t)
	checkFloat64(d.Sum(), 2.3, TOL, "Sum", t)
	checkFloat64(d.Mean(), 2.3, TOL, "Mean", t)
	checkNaN(d.PopulationVariance(), "PopulationVariance", t)
	checkNaN(d.SampleVariance(), "SampleVariance", t)
	checkNaN(d.PopulationStandardDeviation(), "PopulationStandardDeviation", t)
	checkNaN(d.SampleStandardDeviation(), "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 2 values
func TestAppend2(t *testing.T) {
	var d Desc
	d.Append(2.3)
	d.Append(0.4)
	checkInt(d.Count(), 2, "Count", t)
	checkFloat64(d.Min(), 0.4, TOL, "Min", t)
	checkFloat64(d.Max(), 2.3, TOL, "Max", t)
	checkFloat64(d.Sum(), 2.7, TOL, "Sum", t)
	checkFloat64(d.Mean(), 1.35, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 0.9025, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 1.805, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 0.95, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 1.34350288425444, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), 0.0, TOL, "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), -2.0, TOL, "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 3 values. 
func TestAppend3(t *testing.T) {
	var d Desc
	d.Append(2.3)
	d.Append(0.4)
	d.Append(-3.4)
	checkInt(d.Count(), 3, "Count", t)
	checkFloat64(d.Min(), -3.4, TOL, "Min", t)
	checkFloat64(d.Max(), 2.3, TOL, "Max", t)
	checkFloat64(d.Sum(), -0.7, TOL, "Sum", t)
	checkFloat64(d.Mean(), -0.2333333333333334, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 5.615555555555554, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 8.42333333333333, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 2.36971634495683, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 2.90229794013870, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), -0.3818017741606063, TOL, "PopulationSkew", t)
	checkFloat64(d.SampleSkew(), -0.935219529582825, TOL, "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), -1.5, TOL, "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 4 values. Now all of the statistics are available.
func TestAppend4(t *testing.T) {
	var d Desc
	d.Append(1.0)
	d.Append(2.0)
	d.Append(3.0)
	d.Append(4.0)
	checkInt(d.Count(), 4, "Count", t)
	checkFloat64(d.Min(), 1.0, TOL, "Min", t)
	checkFloat64(d.Max(), 4.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 10.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 2.5, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 1.25, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 1.666666666666667, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 1.118033988749895, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 1.290994448735806, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), 0.0, TOL, "PopulationSkew", t)
	checkFloat64(d.SampleSkew(), 0.0, TOL, "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), -1.36, TOL, "PopulationKurtosis", t)
	checkFloat64(d.SampleKurtosis(), -1.2, TOL, "SampleKurtosis", t)
}

func TestAppend5(t *testing.T) {
	var d Desc
	d.Append(1.0)
	d.Append(2.0)
	d.Append(3.0)
	d.Append(4.0)
	d.Append(5.0)
	checkInt(d.Count(), 5, "Count", t)
	checkFloat64(d.Min(), 1.0, TOL, "Min", t)
	checkFloat64(d.Max(), 5.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 15.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 3.0, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 2.0, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 2.5, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 1.414213562373095, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 1.5811388300841898, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), 0.0, TOL, "PopulationSkew", t)
	checkFloat64(d.SampleSkew(), 0.0, TOL, "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), -1.3, TOL, "PopulationKurtosis", t)
	checkFloat64(d.SampleKurtosis(), -1.2, TOL, "SampleKurtosis", t)
}

func TestAppend10(t *testing.T) {
	var d Desc
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	for _, v := range a {
		d.Append(v)
	}
	checkInt(d.Count(), 10, "Count", t)
	checkFloat64(d.Min(), -123.4, TOL, "Min", t)
	checkFloat64(d.Max(), 115.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 62.83, TOL, "Sum", t)
	checkFloat64(d.Mean(), 6.283, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 3165.19316100, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 3516.88129, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 56.2600494223032, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 59.3032991493728, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), -0.4770396201629045, TOL, "PopulationSkew", t)
	checkFloat64(d.SampleSkew(), -0.565699400196136, TOL, "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), 1.253240236214162, TOL, "PopulationKurtosis", t)
	checkFloat64(d.SampleKurtosis(), 3.179835417592894, TOL, "SampleKurtosis", t)
}

// Append by array. In this case, we use slices to update via half of the array at a time.
func TestAppendArray10(t *testing.T) {
	var d Desc
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	// load the first half of the array
	d.AppendArray(a[:5])
	checkInt(d.Count(), 5, "Count", t)
	checkFloat64(d.Min(), -2.0, TOL, "Min", t)
	checkFloat64(d.Max(), 115.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 174, TOL, "Sum", t)
	checkFloat64(d.Mean(), 34.8, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 1910.56, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 2388.2, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 43.70995309995196, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 48.8692132124101, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), 1.003118841855798, TOL, "PopulationSkew", t)
	checkFloat64(d.SampleSkew(), 1.495361279933617, TOL, "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), -0.5476524250400354, TOL, "PopulationKurtosis", t)
	checkFloat64(d.SampleKurtosis(), 1.809390299839858, TOL, "SampleKurtosis", t)

	// load rest of array. The results will be cumulative.
	d.AppendArray(a[5:])
	checkInt(d.Count(), 10, "Count", t)
	checkFloat64(d.Min(), -123.4, TOL, "Min", t)
	checkFloat64(d.Max(), 115.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 62.83, TOL, "Sum", t)
	checkFloat64(d.Mean(), 6.283, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 3165.19316100, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 3516.88129, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 56.2600494223032, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 59.3032991493728, TOL, "SampleStandardDeviation", t)
	checkFloat64(d.PopulationSkew(), -0.4770396201629045, TOL, "PopulationSkew", t)
	checkFloat64(d.SampleSkew(), -0.565699400196136, TOL, "SampleSkew", t)
	checkFloat64(d.PopulationKurtosis(), 1.253240236214162, TOL, "PopulationKurtosis", t)
	checkFloat64(d.SampleKurtosis(), 3.179835417592894, TOL, "SampleKurtosis", t)
}

// Test the batch functions. Calculate the descriptive stats on the whole array.
func TestArrayStats(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	checkInt(StatsCount(a), 5, "Count", t)
	checkFloat64(StatsMin(a), 1.0, TOL, "Min", t)
	checkFloat64(StatsMax(a), 5.0, TOL, "Max", t)
	checkFloat64(StatsSum(a), 15.0, TOL, "Sum", t)
	checkFloat64(StatsMean(a), 3.0, TOL, "Mean", t)
	checkFloat64(StatsPopulationVariance(a), 2.0, TOL, "PopulationVariance", t)
	checkFloat64(StatsSampleVariance(a), 2.5, TOL, "SampleVariance", t)
	checkFloat64(StatsPopulationStandardDeviation(a), 1.414213562373095, TOL, "PopulationStandardDeviation", t)
	checkFloat64(StatsSampleStandardDeviation(a), 1.5811388300841898, TOL, "SampleStandardDeviation", t)
	checkFloat64(StatsPopulationSkew(a), 0.0, TOL, "PopulationSkew", t)
	checkFloat64(StatsSampleSkew(a), 0.0, TOL, "SampleSkew", t)
	checkFloat64(StatsPopulationKurtosis(a), -1.3, TOL, "PopulationKurtosis", t)
	checkFloat64(StatsSampleKurtosis(a), -1.2, TOL, "SampleKurtosis", t)
}


func TestArrayStats2(t *testing.T) {
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	checkInt(StatsCount(a), 10, "Count", t)
	checkFloat64(StatsMin(a), -123.4, TOL, "Min", t)
	checkFloat64(StatsMax(a), 115.0, TOL, "Max", t)
	checkFloat64(StatsSum(a), 62.83, TOL, "Sum", t)
	checkFloat64(StatsMean(a), 6.283, TOL, "Mean", t)
	checkFloat64(StatsPopulationVariance(a), 3165.19316100, TOL, "PopulationVariance", t)
	checkFloat64(StatsSampleVariance(a), 3516.88129, TOL, "SampleVariance", t)
	checkFloat64(StatsPopulationStandardDeviation(a), 56.2600494223032, TOL, "PopulationStandardDeviation", t)
	checkFloat64(StatsSampleStandardDeviation(a), 59.3032991493728, TOL, "SampleStandardDeviation", t)
	checkFloat64(StatsPopulationSkew(a), -0.4770396201629045, TOL, "PopulationSkew", t)
	checkFloat64(StatsSampleSkew(a), -0.565699400196136, TOL, "SampleSkew", t)
	checkFloat64(StatsPopulationKurtosis(a), 1.253240236214162, TOL, "PopulationKurtosis", t)
	checkFloat64(StatsSampleKurtosis(a), 3.179835417592894, TOL, "SampleKurtosis", t)
}

//
//
// Benchmark tests
//
// run with: gotest -bench="Benchmark"
//

func BenchmarkAppend(b *testing.B) {
	var d Desc
	for i := 0; i < b.N; i++ {
		d.Append(3.5)
	}
}

// Test the incremental Variance function by itself. This result is how fast the 
// Variance is calculated not including the time to incrementally update the Desc
// structure with 10 values.
func BenchmarkPopulationVariance10(b *testing.B) {
	b.StopTimer()
	var d Desc
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	for _, v := range a {
		d.Append(v)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.PopulationVariance()
	}
}

// Test the incremental Variance function by itself. This result is how fast the 
// Variance is calculated _including_ the time to incrementally update the Desc
// structure with 10 values. Therefore this result can be compared to the CalcVariance
// function operating on 10 values.
func BenchmarkPopulationVarWUpdates10(b *testing.B) {
	b.StopTimer()
	var d Desc
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	for _, v := range a {
		d.Append(v)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.PopulationVariance()
	}
}

// benchmark on 10 values, so divide by 10 to estimate the per-value time for array calculations
func BenchmarkCalcPopulationVariance10(b *testing.B) {
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	for i := 0; i < b.N; i++ {
		StatsPopulationVariance(a)
	}
}

func BenchmarkCalcPopulationKurtosis10(b *testing.B) {
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	for i := 0; i < b.N; i++ {
		StatsPopulationKurtosis(a)
	}
}

func BenchmarkCalcSampleKurtosis10(b *testing.B) {
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0, -0.03, -123.4, 23.0, -23.04, 12.3}
	for i := 0; i < b.N; i++ {
		StatsSampleKurtosis(a)
	}
}

// Find the time to calculate Sample Kurtosis on an input array 100k random values.
// The benchmark will repeat this test b.N times to determine a stable time. The 
// resulting stable time is the time for the calculation on 100k values.
func BenchmarkCalcSampleKurtosis100k(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Nanoseconds())
	n := 100000 // not the same as b.N
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Float64()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		StatsSampleKurtosis(a)
	}
}

//
//
// Degenerate examples tests
//
//

// Append() 1 0 value
func TestAppend01(t *testing.T) {
	var d Desc
	d.Append(0.0)
	checkInt(d.Count(), 1, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkNaN(d.PopulationVariance(), "PopulationVariance", t)
	checkNaN(d.SampleVariance(), "SampleVariance", t)
	checkNaN(d.PopulationStandardDeviation(), "PopulationStandardDeviation", t)
	checkNaN(d.SampleStandardDeviation(), "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 2 0 values
func TestAppend02(t *testing.T) {
	var d Desc
	d.Append(0.0)
	d.Append(0.0)
	checkInt(d.Count(), 2, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 0.0, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 0.0, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 0.0, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 0.0, TOL, "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 3 0 values. 
func TestAppend03(t *testing.T) {
	var d Desc
	d.Append(0.0)
	d.Append(0.0)
	d.Append(0.0)
	checkInt(d.Count(), 3, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 0.0, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 0.0, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 0.0, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 0.0, TOL, "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

// Append() 4 0 values. 
func TestAppend04(t *testing.T) {
	var d Desc
	d.Append(0.0)
	d.Append(0.0)
	d.Append(0.0)
	d.Append(0.0)
	checkInt(d.Count(), 4, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 0.0, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 0.0, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 0.0, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 0.0, TOL, "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

func TestAppend05(t *testing.T) {
	var d Desc
	d.Append(0.0)
	d.Append(0.0)
	d.Append(0.0)
	d.Append(0.0)
	d.Append(0.0)
	checkInt(d.Count(), 5, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 0.0, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 0.0, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 0.0, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 0.0, TOL, "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}

func TestAppend010(t *testing.T) {
	var d Desc
	a := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	for _, v := range a {
		d.Append(v)
	}
	checkInt(d.Count(), 10, "Count", t)
	checkFloat64(d.Min(), 0.0, TOL, "Min", t)
	checkFloat64(d.Max(), 0.0, TOL, "Max", t)
	checkFloat64(d.Sum(), 0.0, TOL, "Sum", t)
	checkFloat64(d.Mean(), 0.0, TOL, "Mean", t)
	checkFloat64(d.PopulationVariance(), 0.0, TOL, "PopulationVariance", t)
	checkFloat64(d.SampleVariance(), 0.0, TOL, "SampleVariance", t)
	checkFloat64(d.PopulationStandardDeviation(), 0.0, TOL, "PopulationStandardDeviation", t)
	checkFloat64(d.SampleStandardDeviation(), 0.0, TOL, "SampleStandardDeviation", t)
	checkNaN(d.PopulationSkew(), "PopulationSkew", t)
	checkNaN(d.SampleSkew(), "SampleSkew", t)
	checkNaN(d.PopulationKurtosis(), "PopulationKurtosis", t)
	checkNaN(d.SampleKurtosis(), "SampleKurtosis", t)
}





//
//
// Assertion functions used for tests
//
//

// check that the value x equals the expected value y
func checkInt(x, y int, test string, t *testing.T) {
	if x != y {
		t.Errorf("Found %v, but expected %v for test %v", x, y, test)
	}
}

func checkFloat64(x, y, tol float64, test string, t *testing.T) {
	if math.Fabs(x-y) > math.Fabs(x*tol) {
		t.Errorf("Found %v, but expected %v for test %v", x, y, test)
	}
}

func checkNaN(x float64, test string, t *testing.T) {
	if !math.IsNaN(x) {
		t.Errorf("Found %v, but expected NaN for test %v", x, test)
	}
}

func checkFloat64Abs(x, y, tol float64, test string, t *testing.T) {
	if math.Fabs(x-y) > math.Fabs(tol) {
		t.Errorf("Found %v, but expected %v for test %v", x, y, test)
	}
}

func checkInf(x float64, test string, t *testing.T) {
	if !math.IsInf(x, 1) {
		t.Errorf("Found %v, but expected Inf for test %v", x, test)
	}
}
