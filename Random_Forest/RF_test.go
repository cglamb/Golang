package main

import (
	"reflect"
	"testing"

	"github.com/fxsjy/RF.go/RF"
)

// test for ConvertStringSliceToInterface
func TestConvertStringSliceToInterface_SpecificCase(t *testing.T) {
	input := [][]string{{"0", "1"}, {"1", "0"}}
	expected := [][]interface{}{{"0", "1"}, {"1", "0"}}

	result := ConvertStringSliceToInterface(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

// test for ConvertDFYToStringSlice
func TestConvertDFYToStringSlice(t *testing.T) {
	input := [][]string{{"0"}, {"0"}, {"1"}}
	expected := []string{"0", "0", "1"}

	result := ConvertDFYToStringSlice(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// test metrics calculations
func TestCalcMetrics(t *testing.T) {
	predictions := []string{"0", "1", "1", "0"}
	actuals := []string{"0", "1", "0", "0"}
	expectedAccuracy := 0.75
	expectedPrecision := 0.5
	expectedRecall := 1.0
	expectedf1 := 0.6666666666666666

	accuracy, precision, recall, f1 := calcMetrics(predictions, actuals)

	if accuracy != expectedAccuracy {
		t.Errorf("Expected accuracy of %v, got %v", expectedAccuracy, accuracy)
	}

	if precision != expectedPrecision {
		t.Errorf("Expected precision of %v, got %v", expectedPrecision, precision)
	}

	if recall != expectedRecall {
		t.Errorf("Expected recall of %v, got %v", expectedRecall, recall)
	}

	if f1 != expectedf1 {
		t.Errorf("Expected recall of %v, got %v", expectedf1, f1)
	}
}

//unit testing for the random forest model will be completed via checking model fit against parrallel model ran in R and python

// benchmark for model fitting
func BenchmarkEntireProcess(b *testing.B) {

	x_train, _ := processXData("x_train.csv", "Data")
	y_train, _ := processYData("y_train.csv", "Data")
	x_test, _ := processXData("x_test.csv", "Data")

	for i := 0; i < b.N; i++ {

		forest := RF.BuildForest(x_train, y_train, 10, len(x_train[0]), 10)
		// forest := RF.BuildForest(x_train, y_train, 100, len(x_train[0]), 10)
		// forest := RF.BuildForest(x_train, y_train, 100, len(x_train[0]), 50)
		// forest := RF.BuildForest(x_train, y_train, 100, len(x_train[0]), 100)
		// forest := RF.BuildForest(x_train, y_train, 1000, len(x_train[0]), 10)

		for _, testRow := range x_test {
			_ = forest.Predicate(testRow)
		}
	}
}

// // benchmark for data processing
// func BenchmarkProcessXData(b *testing.B) {

// 	fileName := "x_train.csv"
// 	filename2 := "x_test.csv"
// 	subFolder := "Data"

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		_, err := processXData(fileName, subFolder)
// 		if err != nil {
// 			b.Fatalf("Error processing X data")
// 		}
// 		_, err2 := processXData(filename2, subFolder)
// 		if err2 != nil {
// 			b.Fatalf("Error processing X data")
// 		}
// 	}
// }

// func BenchmarkProcessYData(b *testing.B) {

// 	fileName := "y_train.csv"
// 	fileName2 := "y_test.csv"
// 	subFolder := "Data"

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		_, err := processYData(fileName, subFolder)
// 		if err != nil {
// 			b.Fatalf("Error processing Y data: %v", err)
// 		}
// 		_, err2 := processYData(fileName2, subFolder)
// 		if err2 != nil {
// 			b.Fatalf("Error processing Y data: %v", err)
// 		}
// 	}
// }

// func BenchmarkRandomForestBuild(b *testing.B) {

// 	x_train, _ := processXData("x_train.csv", "Data")
// 	y_train, _ := processYData("y_train.csv", "Data")
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = RF.BuildForest(x_train, y_train, 10, 500, len(x_train[0]))
// 	}
// }

// func BenchmarkMakePredictions(b *testing.B) {
// 	x_train, _ := processXData("x_train.csv", "Data")
// 	y_train, _ := processYData("y_train.csv", "Data")
// 	x_test, _ := processXData("x_test.csv", "Data")
// 	forest := RF.BuildForest(x_train, y_train, 10, 500, len(x_train[0]))
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		for _, testRow := range x_test {
// 			_ = forest.Predicate(testRow)
// 		}
// 	}
// }
