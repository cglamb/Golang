package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

// Fits a line using the montanaflynn package linear regression function
// Uses output of .LinearRegression to calculate intercept and slope
func regress(data []stats.Coordinate) (float64, float64) {
	r, _ := stats.LinearRegression(data)

	var X_0, X_1, Y_0, Y_1 float64

	//loop through to find two items in series with different x values
	//once we find different x values also stores associated y values
	for i := 1; i < len(r); i++ {
		if r[i].X != r[0].X {
			X_0 = r[0].X
			X_1 = r[i].X
			Y_0 = r[0].Y
			Y_1 = r[i].Y
			break
		}
	}

	//calculates intercept and slope based on the two points from the series above
	slope := (Y_1 - Y_0) / (X_1 - X_0)
	intercept := Y_0 - slope*X_0
	return slope, intercept
}

// Print results of linear regression
func print_coeff(datasetName string, slope, intercept float64) {
	fmt.Printf("%s:\n", datasetName)
	fmt.Printf("Slope of the regression line: %.4f\n", slope)
	fmt.Printf("Intercept of the regression line: %.4f\n", intercept)
}

// Regress on the 4 Anscombe datasets
func main() {
	// Dataset 1
	data1 := []stats.Coordinate{
		{10, 8.04},
		{8, 6.95},
		{13, 7.58},
		{9, 8.81},
		{11, 8.33},
		{14, 9.96},
		{6, 7.24},
		{4, 4.26},
		{12, 10.84},
		{7, 4.82},
		{5, 5.68},
	}

	slope1, intercept1 := regress(data1)
	print_coeff("Data set 1", slope1, intercept1)

	// Dataset 2
	data2 := []stats.Coordinate{
		{10, 9.14},
		{8, 8.14},
		{13, 8.74},
		{9, 8.77},
		{11, 9.26},
		{14, 8.10},
		{6, 6.13},
		{4, 3.10},
		{12, 9.13},
		{7, 7.26},
		{5, 4.74},
	}

	slope2, intercept2 := regress(data2)
	print_coeff("Data set 2", slope2, intercept2)

	// Dataset 3
	data3 := []stats.Coordinate{
		{10, 7.46},
		{8, 6.77},
		{13, 12.74},
		{9, 7.11},
		{11, 7.81},
		{14, 8.84},
		{6, 6.08},
		{4, 5.39},
		{12, 8.15},
		{7, 6.42},
		{5, 5.73},
	}

	slope3, intercept3 := regress(data3)
	print_coeff("Data set 3", slope3, intercept3)

	// Dataset 4
	data4 := []stats.Coordinate{
		{8, 6.58},
		{8, 5.76},
		{8, 7.71},
		{8, 8.84},
		{8, 8.47},
		{8, 7.04},
		{8, 5.25},
		{19, 12.5},
		{8, 5.56},
		{8, 7.91},
		{8, 6.89},
	}

	slope4, intercept4 := regress(data4)
	print_coeff("Data set 4", slope4, intercept4)

}
