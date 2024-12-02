package main

import (
	"log"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// converts to depedent variable to a []float64 slice
// dependent variable is in a gota dataframe
// we use this function to convert to type we can do matrix math on
func depedentSlicer(records [][]string) []float64 {

	data := make([]float64, len(records)-1) //creating a data slice to store records.  excluding heater so 1 row shorter than input data
	for i, record := range records[1:] {    //iterating over all records but the header row
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatalf("error parsing dependent variable.  error in depedentSlier.  a record may not be of a type that can be coverted to float64")
		}
		data[i] = value
	}
	return data
}

// convert the indepedent variables in the dataframe to a [][]float64
// data is in a gota dataframe
// we use this function to convert to type we can do matrix math on
func indepdentSlicer(records [][]string) [][]float64 {

	data := make([][]float64, len(records)-1) //slice of slice to store indepedent variable data.  excluding header so 1 row short than input data
	for i, record := range records[1:] {      //iteration over everything but the header
		row := make([]float64, len(record)) //row slice to append back into data
		for j, cell := range record {
			val, err := strconv.ParseFloat(cell, 64)
			if err != nil {
				log.Fatalf("error parsing indepedent variable.  error in indepdentSlicer.  a record may not be of a type that can be coverted to float64")
			}
			row[j] = val
		}
		data[i] = row
	}
	return data
}

func createLevels(df dataframe.DataFrame, columnName string) dataframe.DataFrame {

	// pop the columing being leveled
	targetCol := df.Col(columnName)

	// create a list of all unique levels
	levelNames := targetCol.Records()  //grab all the records
	uniqueMap := make(map[string]int)  //initalize a map
	var levels []string                //slice that will contain the level name
	for _, level := range levelNames { //iterate through all the records.  save level name if not already in lsit
		if _, exists := uniqueMap[level]; !exists {
			uniqueMap[level] = len(levels)
			levels = append(levels, level)
		}
	}

	// index the level names
	levelsSlice := make([]int, len(levelNames)) //initialize slice for level index
	for i, level := range levelNames {          //looping through names and assinging a level index
		levelsSlice[i] = uniqueMap[level]
	}

	// drop old column and replace with level
	df = df.Drop([]string{"neighborhood"})
	levelsSeries := series.New(levelsSlice, series.Int, columnName)
	df = df.CBind(dataframe.New(levelsSeries))

	return df
}
