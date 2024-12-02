package main

//package executedin terminal via go test -bench=/ command

import (
	"testing"
)

// i cant get the types to match up here.  needs more work
// func TestProcessData(t *testing.T) {
// 	data, _ := readCSV("testdata/housesInput.csv")
// 	result := processData(data)

// 	expected := [][]interface{}{
// 		{"value", 20640, 206855.81690891474, 115395.6158744132, 14999, 119600, 179700, 264725, 500001},
// 		{"income", 20640, 3.8706710029070246, 1.8998217179452732, 0.4999, 2.5633999999999997, 3.5347999999999997, 4.74325, 15.0001},
// 		{"age", 20640, 28.639486434108527, 12.585557612111637, 1, 18, 29, 37, 52},
// 		{"rooms", 20640, 2635.7630813953488, 2181.615251582787, 2, 1447.75, 2127, 3148, 39320},
// 		{"bedrooms", 20640, 537.8980135658915, 421.2479059431317, 1, 295, 435, 647, 6445},
// 		{"pop", 20640, 1425.4767441860465, 1132.4621217653375, 3, 787, 1166, 1725, 35682},
// 		{"hh", 20640, 499.5396802325581, 382.3297528316099, 1, 280, 409, 605, 6082},
// 	}

// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Result %v is not equal to expected %v", result, expected)
// 	}

// }

func BenchmarkProcessDataLoop(b *testing.B) {

	data, _ := readCSV("testdata/housesInput.csv")

	b.ResetTimer()

	_ = processDataLoop(data, 100)

}
