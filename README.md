
# Stats #

Stats is a descriptive statistics and linear regression package for Go. It provides:

* Descriptive Statistics: count, sum, min, max, mean, variance, standard deviation, skew, and kurtosis
* Linear Regression: slope, intercept, r-squared, slope standard error, intercept standard error
* Incremental updates: the stats and regression can be updated one or a few at a time.
* Batch updates: Calculate stats and regression only for the given array of values.
* Population and sample statistics included

Incremental updates are useful for streaming data applications or situations in which storing the data for statistics is prohibitive. In fact, if the data is stored only for the purpose of statistical calculations, incremental updates make storage unnecessary.

The package includes convenience functions that allow incremental updates by single or multiple values. Or you can use traditional batch calculations on a given array of values. The linear regression functions also include incremental and batch updates. 

See demo.go and the *_test.go files for example usage.

## Installation

To install

	git clone https://github.com/GaryBoone/GoStats.git
	cd GoStats
	make
	make install

To make and run the demo

	cd demo
	make
	./demo

## Usage



## Tests

	gotest
	
## Benchmarks

	gotest -bench="Benchmark"