package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// structure that will hold each iteration of our regression model
type dfResults struct {
	Coeffs []string
	MSE    float64
	AIC    float64
}

// bundles the regression calculation and associated regression calculations: mse and AIC
func regressionBundle(dfMatrix [][]float64, headers []string, targetCombo []string, mvSlice []float64) ([]string, float64, float64, error) {

	//creates a modeling database that contains only the data from the indepedent variables relevant to this regressoin
	//that is to say this is a subset of the full dataset of indepdent variables
	dfModeling, err := modelingData(dfMatrix, headers, targetCombo)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error generating modeling database: %v", err)
	}

	//run the OLS regression
	coeffs, err := LinearRegressionFit(dfModeling, mvSlice)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error performing the OLS regression: %v", err)
	}

	//calculated the predicted yhat for each x in the dataset
	predictions := Predict(dfModeling, coeffs)

	//run the MSE calculation
	mse := meanSquareError(mvSlice, predictions)

	// Calculating AIC
	nobs := float64(len(mvSlice))
	ssr := mse * nobs
	llf := -nobs/2*math.Log(2*math.Pi) - nobs/2*math.Log(ssr/nobs) - nobs/2
	aic := -2*llf + 2*float64(len(coeffs))

	return targetCombo, mse, aic, nil
}

// outputs best model
// best model measured by lowest AIC
// did not address ties -- future area of enchancement
func bestModel(results []dfResults) ([]string, float64, float64) {

	besti := 0
	minAIC := results[0].AIC

	// find the model with the lowest AI
	for i, result := range results {
		if result.AIC < minAIC {
			besti = i
			minAIC = result.AIC
		}
	}

	// returns info on the best model
	bestModel := results[besti]
	return bestModel.Coeffs, bestModel.MSE, bestModel.AIC
}

// run the regressions without concurrency...ie sequentially
// n_times is the number of duplicate times we will run each regression
// clean combos is all the possible combinations of independent variables we will test
// dfMatrix is the dataset containing all independent variables
// headers are the labels for all indpendent variables
// mvslice is the depedent variable
func runSeq(n_times int, clean_combos [][]string, dfMatrix [][]float64, headers []string, mvSlice []float64) {
	var results []dfResults        //will save aggregate results
	for x := 0; x < n_times; x++ { //looping to do each regresion calculation n_times...this is used for benchmarking
		for _, i := range clean_combos { //looping through each combo of indepedent variables
			coeffs, mse, aic, err := regressionBundle(dfMatrix, headers, i, mvSlice) //running the regression and calculating relevant statistics
			if err != nil {
				log.Fatalf("error running regression.  error in function runSeq")
			}
			results = append(results, dfResults{Coeffs: coeffs, MSE: mse, AIC: aic}) //add this combos results to full data set
		}
	}
	n_models := len(results)               //calculate number of total regressions we are doing
	coeffs, mse, aic := bestModel(results) //searches database to find model with lowest AIC.  thats selected as our best model
	fmt.Println("---Running regresions sequentially---")
	fmt.Println("Number of models run:", n_models)
	fmt.Printf("Best Model Beta Coefficients: %v\nMSE: %f\nAIC: %f\n", coeffs, mse, aic)
}

// adds concurrency to runSeq function
// function is identical to sequesntial model other than for concurrency
// code comments are included only for concurrency
// see runSeq() for any comments associated with non-concurrent elements
func runConcur(n_times int, clean_combos [][]string, dfMatrix [][]float64, headers []string, mvSlice []float64) {

	var wg sync.WaitGroup                                             //creating the struct for the goroutine coordination
	resultsChannel := make(chan dfResults, len(clean_combos)*n_times) //channel to collect the regression results

	for x := 0; x < n_times; x++ {
		for _, combo := range clean_combos {
			wg.Add(1)                 //add task to que
			go func(combo []string) { //passing to the new gouroutine
				defer wg.Done() //free up resource
				coeffs, mse, aic, err := regressionBundle(dfMatrix, headers, combo, mvSlice)
				if err != nil {
					log.Printf("error in regression bundle: %v", err)
					return
				}
				resultsChannel <- dfResults{Coeffs: coeffs, MSE: mse, AIC: aic}
			}(combo)
		}
	}

	wg.Wait() //wait for all goroutines to complete
	close(resultsChannel)

	var results []dfResults
	for result := range resultsChannel {
		results = append(results, result)
	}

	n_models := len(results)
	coeffs, mse, aic := bestModel(results)
	fmt.Println("---Running regresions concurrently---")
	fmt.Println("Number of models run:", n_models)
	fmt.Printf("Best Model Beta Coefficients: %v\nMSE: %f\nAIC: %f\n", coeffs, mse, aic)
}

func main() {

	input_filename := "boston.csv"
	// n_times := 1

	var n_times int
	var useCon bool

	// user input for command line application
	// ask user for number of times to run each regression
	flag.IntVar(&n_times, "i", 1, "number of times to loop through the calculation")
	flag.BoolVar(&useCon, "c", true, "run with concurrency")
	flag.Parse()

	// Open the CSV file.
	csvFile, _ := ReadCSV(input_filename)
	defer csvFile.Close()

	// put the csv into a dataframe
	// im using a gota datframe to increase the ease of data manipulation
	df := dataframe.ReadCSV(csvFile)

	//begin some data cleansing

	// creating levels for neighboorhood
	df = createLevels(df, "neighborhood")
	//converting chas, rad, and tax to float
	chas_new := df.Col("chas").Float()
	rad_new := df.Col("rad").Float()
	tax_new := df.Col("tax").Float()
	neighbhorhood_new := df.Col("neighborhood").Float()
	df = df.Drop([]string{"chas", "rad", "tax", "neighborhood"})
	chasNewDF := dataframe.New(series.New(chas_new, series.Float, "chas"))
	radNewDF := dataframe.New(series.New(rad_new, series.Float, "rad"))
	taxNewDF := dataframe.New(series.New(tax_new, series.Float, "tax"))
	neighbhorhoodNewDF := dataframe.New(series.New(neighbhorhood_new, series.Float, "neighborhood"))
	df = df.CBind(chasNewDF).CBind(radNewDF).CBind(taxNewDF).CBind(neighbhorhoodNewDF)

	//popping out mv to save as seperate dependent variable
	mv := df.Select([]string{"mv"})
	df = df.Drop([]string{"mv"})

	//converting mv to []float64
	mvRecords := mv.Records()
	mvSlice := depedentSlicer(mvRecords)

	//convert indepdent variables to [][]float64
	records := df.Records()
	dfMatrix := indepdentSlicer(records)

	//generating an exhuastive list of indepdent variable combinations
	//this contains a subset of columns in dfMatrix
	//we will run a linear regression on each combination in this slice
	headers := df.Names()                              //grab all headers
	combo_list := exhaustiveList(headers)              //creates a list of all possible combination of independent variables
	clean_combos := filteredExhaustiveList(combo_list) //removes any empty strings

	// run either with concurrency or sequentially
	if useCon {
		runConcur(n_times, clean_combos, dfMatrix, headers, mvSlice)
	} else {
		runSeq(n_times, clean_combos, dfMatrix, headers, mvSlice)
	}

}
