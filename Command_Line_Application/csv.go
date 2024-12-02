package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV Reader used to read in data from CSV files
	cr := csv.NewReader(r)

	// Adjusting for 0 based index
	column--

	// Read in all CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Cannot read data from file: %w", err)
	}

	var data []float64

	// Looping through all records
	for i, row := range allData {
		if i == 0 {
			continue
		}

		// Checking number of columns in CSV file
		if len(row) <= column {
			// File does not have that many columns
			return nil,
				fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}

		// Try to convert data read into a float number
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	// Return the slice of float64 and nil error
	return data, nil
}

// read in CSV
func readCSV(filePath string) ([][]string, error) {

	// open path
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the csv
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	//output results
	return records, nil
}

// takes data from csv and returns only headers
func readHeaders(results [][]string) []string {
	if len(results) > 0 {
		return results[0]
	}
	return nil
}

// takes data from csv and returns only data elements
// removes headers
func readData(results [][]string) [][]string {
	if len(results) > 1 {
		return results[1:]
	}
	return nil
}

// convert string data to float64
func convertFloat(data [][]string) [][]float64 {

	var floatData [][]float64

	for _, row := range data {
		var floatRow []float64
		for _, value := range row {
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				// Handle conversion error if necessary
				fmt.Println("Error converting to float:", err)
				continue
			}
			floatRow = append(floatRow, floatValue)
		}
		floatData = append(floatData, floatRow)
	}

	return floatData
}

// return the nth column in the csv data
func readCol(floatData [][]float64, n int) []float64 {
	var result []float64
	for _, row := range floatData {
		if n < len(row) {
			result = append(result, row[n])
		}
	}
	return result
}

// write to txt
func writeTxt(filename string, data []interface{}) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, item := range data {
		row, ok := item.([]interface{})
		if !ok {
			return errors.New("Invalid data format in universalSlice")
		}

		for _, value := range row {
			fmt.Fprintf(outputFile, "%v ", value)
		}
		fmt.Fprintln(outputFile)
	}

	return nil
}
