// fits a random forest model to the training data and then uses the model to make predictions on the test data
// accuracy, precision, and recall are then calculated and printed to the console
// random forest library based on:
// https://github.com/fxsjy/RF.go/blob/master/RF/Forest.go

package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/fxsjy/RF.go/RF"
)

// readCSV reads a csv file and returns a slice of slices
func readCSV(fileName string, subfolder string) ([][]string, error) {

	filePath := filepath.Join(subfolder, fileName) //create the file path

	file, err := os.Open(filePath) //open the file
	if err != nil {
		log.Fatalf("Unable to open file %s: %v", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file) //create a new reader

	//read the file into a slice of slices
	df_temp, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", filePath, err)
	}

	//strip the headers from the data to create test training data
	df := df_temp[1:] //strip the headers from the data

	return df, nil
}

// convert slice of slices of strings to slice of slices of interfaces
// random forset library requires the data to be in an interface
func ConvertStringSliceToInterface(ss [][]string) [][]interface{} {
	var data [][]interface{}
	for _, record := range ss {
		var row []interface{}
		for _, field := range record {
			row = append(row, field) // Convert string to interface{}
		}
		data = append(data, row)
	}
	return data
}

// converts a slice of slices of strings to a slice of strings
// random forset library requires the y values to be in a slice of strings
func ConvertDFYToStringSlice(df_y [][]string) []string {
	var yData []string
	for _, row := range df_y {
		if len(row) > 0 {
			yData = append(yData, row[0])
		}
	}
	return yData
}

// read and output x data
func processXData(fileName string, subfolder string) ([][]interface{}, error) {
	df, err := readCSV(fileName, subfolder)
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", fileName, err)
	}
	xData := ConvertStringSliceToInterface(df)
	return xData, nil
}

// read and output y data
func processYData(fileName string, subfolder string) ([]string, error) {
	df, err := readCSV(fileName, subfolder)
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", fileName, err)
	}
	yData := ConvertDFYToStringSlice(df)
	return yData, nil
}

// calculates accuracy, precision, and recall
func calcMetrics(predictions []string, actuals []string) (float64, float64, float64, float64) {
	TP, TN, FP, FN := confusionMatrix(predictions, actuals)
	accuracy := accruacy(TP, TN, FP, FN)
	precision := precision(TP, FP)
	recall := recall(TP, FN)
	f1 := 2 * ((precision * recall) / (precision + recall))
	return accuracy, precision, recall, f1
}

// accuracy calculation
func accruacy(TP int, TN int, FP int, FN int) float64 {
	return float64(TP+TN) / float64(TP+TN+FP+FN)
}

// precision calculation
func precision(TP int, FP int) float64 {
	return float64(TP) / float64(TP+FP)
}

// recall calculation
func recall(TP int, FN int) float64 {
	return float64(TP) / float64(TP+FN)
}

// calculates TP, TN, FP, FN based on passed predictions and actuals
func confusionMatrix(predictions []string, actuals []string) (int, int, int, int) {

	//intializing zero values for output values
	TP := 0
	TN := 0
	FP := 0
	FN := 0

	//iterating through the predictions and actuals to calculate the confusion matrix
	for i := 0; i < len(predictions); i++ {
		if predictions[i] == "1" {
			if actuals[i] == "1" {
				TP++
			} else {
				FP++
			}
		} else {
			if actuals[i] == "0" {
				TN++
			} else {
				FN++
			}
		}
	}
	return TP, TN, FP, FN
}

func main() {

	// Create a server
	srv := &http.Server{Addr: "localhost:6060"}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	var n_trees int
	var features_at_split int

	flag.IntVar(&n_trees, "n_trees", 100, "number of trees")
	flag.IntVar(&features_at_split, "features_at_split", 100, "features to consider at each split")
	flag.Parse()

	log.Println("starting model training and evaluation")

	//import a csv file containing the training data
	fileName_xtrain := "x_train.csv"
	fileName_ytrain := "y_train.csv"
	fileName_xtest := "x_test.csv"
	fileName_ytest := "y_test.csv"
	subFolder := "Data"

	//read in training data
	log.Printf("reading in data")
	x_train, _ := processXData(fileName_xtrain, subFolder)
	y_train, _ := processYData(fileName_ytrain, subFolder)

	//read in test data
	x_test, _ := processXData(fileName_xtest, subFolder)
	y_test, _ := processYData(fileName_ytest, subFolder)

	// inputs , labels, treesAmount (number of trees), feature count, number of independet varaibles)
	log.Printf("building the random forest model")
	forest := RF.BuildForest(x_train, y_train, n_trees, len(x_train[0]), features_at_split)

	// save a slice of predictions
	log.Printf("making predictions")
	predictions := make([]string, 0, len(x_test)) //initialize an empty slice of strings
	for _, prediction := range x_test {
		prediction := forest.Predicate(prediction)
		predictions = append(predictions, prediction)
	}

	// calculate the fit metrics
	log.Printf("calculating fit metrics")
	accuracy, precision, recall, f1 := calcMetrics(predictions, y_test)
	fmt.Println("Accuracy:", accuracy)
	fmt.Println("Precision:", precision)
	fmt.Println("Recall:", recall)
	fmt.Println("F1:", f1)

	// close the http server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
