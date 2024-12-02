package main

import (
	"log"
	"os"
)

// read in the csv
func ReadCSV(filename string) (*os.File, error) {

	csv, err := os.Open(filename) //open the file
	//generate error reporting
	if err != nil {
		log.Fatal("error reading in csv.  error in readcsv function.  file could not be located or file type invalid") //throw a fatal error if we cant find the file
		return nil, err
	}

	return csv, nil

}
