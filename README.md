
# Stats #

Stats is a descriptive statistics and linear regression package for Go. It provides:

* Descriptive Statistics: count, sum, min, max, mean, variance, standard deviation, skew, and kurtosis
* Univariate Linear Regression: slope, intercept, r-squared, slope standard error, intercept standard error
* Incremental updates: the stats and regression can be updated one or a few at a time.
* Batch updates: Calculate stats and regression only for the given array of values.
* Population and sample statistics included

Incremental updates are useful for streaming data applications or situations in which storing the data for statistics is prohibitive. In fact, if the data is stored only for the purpose of statistical calculations, incremental updates make storage unnecessary.

The package includes convenience functions that allow incremental updates by single or multiple values. Or you can use traditional batch calculations on a given array of values. The linear regression functions also include incremental and batch updates. 


## Installation

Set up Go env as shown in [How to Write Go Code](http://golang.org/doc/code.html) and this [screencast](http://www.youtube.com/watch?v=XCsL89YtqCs).

Install the stats pkg

	$ cd goProjects/src/GoStats/stats
	$ go test -v --bench='.'
	$ go install

Run the descriptive statistics demo

	$ cd ../descriptive_statistics_demo/
	$ go build
	$ ./descriptive_statistics_demo

Run the regression demo

	$ cd ../regression_demo/
	$ go build
	$ ./regression_demo


	

## Usage ##

The GoStats package is flexible and can calculate statistics in _batch_ or _incremental_ modes. If you have an array of values and want to see their descriptive statistics, use the batch methods. They are prefixed with _Stats_, such as _StatsSampleVariance( []float64 )_.

Incremental mode means that you pass values into the GoStats package one or a few at a time. The descriptive statistics are available at any time. This mode is useful in situations in which the data may be generated one or a few at a time by some process. It's also memory efficient because the statistics are available without having to store all of the values they're based on, in contrast to the batch methods.

### Descriptive Statistics ###

#### Batch
Batch updates are the traditional calculations of descriptive statistics on a given array of values. They don't require a Stats struct and are prefixed with 'Stats'.

	a := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	populationVariance := StatsPopulationVariance(a)   // = 2.0
	sampleVariance := StatsSampleVariance(a)           // = 2.5

#### Incremental

To use incremental updates, declare a Stats struct

	var d stats.Stats

Then update it with new values

	d.Update(x)
	
To obtain the descriptive statistics

	count := d.Count()
	min := d.Min()
	max := d.Max()
	sum := d.Sum()
	mean := d.Mean()
	standardDeviation := d.SampleStandardDeviation()
	variance := d.SampleVariance()
	skew := d.SampleSkew()
	kurtosis := d.SampleKurtosis()

Note that population statistics are also provided

	popStandardDeviation := d.PopulationStandardDeviation()
	popVariance := d.PopulationVariance()
	popSkew := d.PopulationSkew()
	popKurtosis := d.PopulationKurtosis()

Updates can also be done with arrays of values

	var d stats.Stats
	a := []float64{1.0, -2.0, 13.0, 47.0, 115.0}
	d.UpdateArray(a)

Note that this is an update to an existing Stats struct. It updates the current values.


	
### Linear Regression ###

Similarly, univariate linear regression can be done incrementally or in batch.

Declare a regression struct to hold the intermediate values

	var r stats.Regression

Then update as data becomes available

	r.Update(x, y)
	
The regression can be calculated at any time and does not affect the Regression struct. So you can continue to update it.
	
	slope := r.Slope()
	intercept := r.Intercept()
	r_squared := r.RSquared()
	count := r.Count()
	slopeStdErr := r.SlopeStandardError()
	interceptStdErr := r.InterceptStandardError()

As before, updates can be given arrays

	var r stats.Regression
	// do some r.Update(x, y) 

	// now update with arrays of values
	xData := []float64{2000, 2001, 2002, 2003, 2004}
	yData := []float64{9.34, 8.50, 7.62, 6.93, 6.60}
	r.UpdateArray(xData, yData)

	// check regression
	slopeStdErr := r.SlopeStandardError()
	interceptStdErr := r.InterceptStandardError()


Batch linear regressions are done by just passing in the x and y arrays
	
	var slope, intercept, rsquared, count, slopeStdErr, intcptStdErr = LinearRegression(xData, yData)

Note that if you don't need all of the values, you can ignore them

	var slope, intercept, _, _, _, _ = LinearRegression(xData, yData)

	
## Tests ##

To test, all code was compared against the [R stats package](http://r-project.org).

	$ go test
	
## Benchmarks ##

The benchmarks show that the incremental and batch functions show similar efficiency.

	$ go test -bench='.'
	
## License ##
The code is available at github [GaryBoone/GoStats](https://github.com/GaryBoone/GoStats) under the [MIT license](http://opensource.org/licenses/mit-license.php).
