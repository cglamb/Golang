package main

import "fmt"

// generate a string that contains all possible header combinations
// order is not considered important
// allows for different combo lengths.  so will have combos with only one independent variable all the way up to all independent variables
// reference: https://blog.enterprisedna.co/how-to-generate-all-combinations-of-a-list-in-python/
func exhaustiveList(headers []string) [][]string {
	result := [][]string{{}} // empty slice of slice to store the header combos
	for _, header := range headers {
		newCombos := [][]string{}      //sub slice to store each unique combo
		for _, combo := range result { //sub-loop to loop through remaining variables
			newCombo := append([]string{}, combo...)
			newCombo = append(newCombo, header)
			newCombos = append(newCombos, newCombo)
		}
		// generate the master list
		result = append(result, newCombos...)
	}
	return result
}

// exhuastiveList includes an empty combo
// this removes the empty combo
// we dont want the empty combo, as we cant run a regression on a dataset with no indepdent variables
func filteredExhaustiveList(results [][]string) [][]string {
	filteredResult := [][]string{}  //empty slice of slices to store ending results
	for _, slice := range results { //loop through each slice
		if len(slice) > 0 { // check if there is something in the slice, ie lenght is greather than 0
			filteredResult = append(filteredResult, slice) //if the slice is populated add it to the filtered list
		}
	}
	return filteredResult
}

// generates a modeling data based on the full indepedent variable data
// takes as input a slice of slices
// a specific combinations of headers is also passed to the function
// the function outputs a subset of the input data that corresponds to only the target independent variables/headers
// we use this function to allow us to exhaustively run the linear regressions
func modelingData(dfMatrix [][]float64, headers, targetCombo []string) ([][]float64, error) {

	columni := make([]int, 0, len(targetCombo)) //dummy index setup
	indexmap := make(map[string]int)            //map assigning each header word to an index number

	//populate the index numbers in the map
	for i, header := range headers {
		indexmap[header] = i
	}

	// search for the index numbers of the headers we want to keep
	for _, header := range targetCombo {
		if index, exists := indexmap[header]; exists {
			columni = append(columni, index)
		} else {
			return nil, fmt.Errorf("eeror finding headers.  see function modelingData")
		}
	}

	// generating the data subset we need
	filteredata := make([][]float64, len(dfMatrix)) // empty string of strings that will eventual contain the output data
	for i, row := range dfMatrix {                  //loop through each row of data in the
		filteredRow := make([]float64, len(columni)) //empty slice for the
		for j, coli := range columni {               //if column index is in list of indepedent variables we are targeting keep
			if coli < len(row) {
				filteredRow[j] = row[coli]
			}
		}
		filteredata[i] = filteredRow
	}

	return filteredata, nil
}
