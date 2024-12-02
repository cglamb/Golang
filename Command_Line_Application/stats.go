package main

import (
	"math"
	"sort"
)

// sum function that we will use for calculating other statistics
func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

// calculate the mean
func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// returns number of elements in the dataset
func count(data []float64) float64 {
	return float64(len(data))
}

// standard deviation calculation
func std(data []float64) float64 {

	// Calculate the mean by calling the average function
	mean := avg(data)

	//sum across the squared deviations
	sq_dev := 0.0

	for _, v := range data {
		sq_dev += math.Pow(v-mean, 2)
	}

	//divide by count and take the square root
	count := count(data)
	return math.Sqrt(sq_dev / (count - 1))
}

// min calculation
func min(data []float64) float64 {

	//Setting the first value in the slice equal to the stored min value
	minS := data[0]

	//looping through all remaining values in the slice and replacing minS with any lower values we find
	for _, v := range data {
		if v < minS {
			minS = v
		}
	}

	return minS

}

// max calculation
func max(data []float64) float64 {

	//Setting the first value in the slice equal to the stored max value
	maxS := data[0]

	//looping through all remaining values in the slice and replacing maxS with any higher values we find
	for _, v := range data {
		if v > maxS {
			maxS = v
		}
	}

	return maxS

}

// linear interpolation function that will be used for decile calculation
func linearInterpolate(data []float64, index float64) float64 {

	// parammeterizing
	x0 := int(index)
	x1 := x0 + 1
	y0 := data[x0]
	y1 := data[x1]
	w := index - float64(x0)

	// solving
	v := (1.0-w)*y0 + w*y1
	return v
}

// calculating the nth percntile
func Percentile(data []float64, percentile float64) float64 {

	// Sort the data in ascending order
	sort.Float64s(data)

	// translating the nth percentile to an index
	n := float64(len(data))
	index := (n - 1) * percentile

	if index == float64(int(index)) {
		// if the index is an integer we can return that value directly
		return data[int(index)]
	} else {
		//else we have to linearly interpolate
		return linearInterpolate(data, index)
	}
}

// returns count, mean, standard deviation, min, max, 25th percentile, 50th percentile, 75th percentile
func calcStats(data []float64) []float64 {

	v_count := count(data)
	v_mean := avg(data)
	v_std := std(data)
	v_min := min(data)
	v_fdec := Percentile(data, .25)
	v_median := Percentile(data, .5)
	v_tdec := Percentile(data, .75)
	v_max := max(data)

	return []float64{v_count, v_mean, v_std, v_min, v_fdec, v_median, v_tdec, v_max}
}
