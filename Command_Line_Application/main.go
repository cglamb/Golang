package main

// acknowledgement - code borrows heavily from:
// Gerardi, Ricardo. 2021. Powerful Command-Line Applications in Go: Build Fast and Maintainable Tools. Raleigh, NC: The Pragmatic Bookshelf.

import (
	"flag"
	"fmt"
	"os"
)

var Master_Slice [][]float64

func main() {

	// collecting user input info
	out_location := flag.String("out_location", "output.txt", "Location to save output")
	filename := flag.String("input_file", "testdata/housesInput.csv", "Location to save output")
	flag.Parse()

	//read in the csv by calling the readCSV function
	csv_data, err := readCSV(*filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	iterations := 100
	data_out := processDataLoop(csv_data, iterations)

	//adding statistical lables and inserting into dataout
	label := []interface{}{"field", "count", "mean", "std", "min", "25%", "50%", "75%", "max"}
	finalout := make([]interface{}, 0, len(data_out)+1)
	finalout = append(finalout, label)
	finalout = append(finalout, data_out...)

	// write the descriptive statistics to txt
	err = writeTxt(*out_location, finalout)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Results have been written to 'output.txt'.")
}

func processData(data [][]string) []interface{} {

	masterSlice := []interface{}{}

	string_data := readData(data)           //removes header from csv data
	float_data := convertFloat(string_data) //converts data to float
	num_cols := len(float_data[0])
	headers := readHeaders(data)

	for i := 1; i <= num_cols; i++ {

		results := make([]interface{}, 0, 10)   //empty inteface to put in stats for each column of data
		results = append(results, headers[i-1]) //inserting the header name as a label

		ithData := readCol(float_data, i-1) //reading the column data
		stats := calcStats(ithData)         //calculating the descriptive statistics

		for _, v := range stats {
			results = append(results, v)
		}

		masterSlice = append(masterSlice, results)
	}

	return masterSlice
}

func processDataLoop(data [][]string, iterations int) []interface{} {

	universalSlice := make([]interface{}, 0)

	for i := 0; i < iterations; i++ {
		result := processData(data)
		universalSlice = append(universalSlice, result...)
	}

	return universalSlice
}
