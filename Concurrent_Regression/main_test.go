package main

import (
	"math"
	"reflect"
	"sort"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// test DependentSlicer
// test reads in a test csv and cheks out values are as expected
func TestDependentSlicer(t *testing.T) {

	test_filename := "test_cases.csv"

	// Open the CSV file.
	csvFile, _ := ReadCSV(test_filename)
	defer csvFile.Close()

	// put the csv into a dataframe
	df_test := dataframe.ReadCSV(csvFile)

	//popping out y to save as seperate dependent variable
	y_df := df_test.Select([]string{"mv"})
	y := y_df.Records()

	// creating the expected results
	expected := []float64{24, 21.6, 34.7}

	//calling the function
	actual := depedentSlicer(y)

	// check the result
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("error")
	}
}

// test indepedent slicer
// test reads in a test csv and cheks out values are as expected
func TestIndepdentSlicer(t *testing.T) {

	test_filename := "test_cases.csv"

	// Open the CSV file.
	csvFile, _ := ReadCSV(test_filename)
	defer csvFile.Close()

	// put the csv into a dataframe
	df_test := dataframe.ReadCSV(csvFile)

	//popping out y to save as seperate dependent variable
	df := df_test.Select([]string{"age", "dis"})
	x := df.Records()

	// creating the expected results
	expected := [][]float64{
		{65.2, 4.09},
		{78.9, 4.9671},
		{61.1, 4.9671},
	}

	//calling the function
	actual := indepdentSlicer(x)

	// check the result
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("error")
	}
}

// testing means square error
func TestMeanSquareError(t *testing.T) {
	y := []float64{1, 2, 3}
	yhat := []float64{1, 5, 3}
	expected := float64(3)

	actual := meanSquareError(y, yhat)

	if actual != expected {
		t.Errorf("error")
	}
}

// testing Linear Regression function
// using sample data only y = 2x + 1
func TestLinearRegressionFit(t *testing.T) {

	//setting up x and ydata
	X := [][]float64{{1}, {2}, {3}, {4}}
	Y := []float64{3, 5, 7, 9}

	// expected coefficients of the OLS regression
	expected := []float64{1, 2}

	// apply the function
	actual, _ := LinearRegressionFit(X, Y)

	// check
	tolerance := 0.001
	for i, b := range actual {
		if !aboutEqual(b, expected[i], tolerance) {
			t.Errorf("error")
		}
	}

}

func aboutEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestExhaustiveList(t *testing.T) {

	headers := []string{"a", "b", "c"}
	expected := [][]string{{}, {"a"}, {"b"}, {"c"}, {"a", "b"}, {"a", "c"}, {"b", "c"}, {"a", "b", "c"}}

	//applying the function
	actual := exhaustiveList(headers)

	// sort slices to avoid differences in order generating test failures
	sortS(actual)
	sortS(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("error")
	}
}

// testing that our filtering function drops any null strings correctly
func TestFilteredExhaustiveList(t *testing.T) {

	//using same data as exhaustive list test
	headers := []string{"a", "b", "c"}
	expected := [][]string{{"a"}, {"b"}, {"c"}, {"a", "b"}, {"a", "c"}, {"b", "c"}, {"a", "b", "c"}}

	//applying the function
	unfiltered := exhaustiveList(headers)
	actual := filteredExhaustiveList(unfiltered)

	// sort slices to avoid differences in order generating test failures
	sortS(actual)
	sortS(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("error")
	}
}

// used for exhaustive list testing
func sortS(slices [][]string) {
	for _, slice := range slices {
		sort.Strings(slice)
	}
	sort.Slice(slices, func(i, j int) bool {
		if len(slices[i]) == 0 {
			return true
		}
		if len(slices[j]) == 0 {
			return false
		}
		return slices[i][0] < slices[j][0]
	})
}

// testing Predict function
func TestPredict(t *testing.T) {

	//lets assume fitted line is y = 2x + 1
	// test at x = 1,2,3
	x := [][]float64{{1}, {2}, {3}}
	beta := []float64{1, 2}
	expected := []float64{3, 5, 7}

	// call
	result := Predict(x, beta)

	// compare
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Predict() = %v; want %v", result, expected)
	}
}

func BenchmarkC(b *testing.B) {

	n_times := 100
	input_filename := "boston.csv"

	//do the prep work to read in and edit the csv
	//this replicates the first part of func main
	csvFile, _ := ReadCSV(input_filename)
	defer csvFile.Close()
	df := dataframe.ReadCSV(csvFile)
	df = createLevels(df, "neighborhood")
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
	mv := df.Select([]string{"mv"})
	df = df.Drop([]string{"mv"})
	mvRecords := mv.Records()
	mvSlice := depedentSlicer(mvRecords)
	records := df.Records()
	dfMatrix := indepdentSlicer(records)
	headers := df.Names()                              //grab all headers
	combo_list := exhaustiveList(headers)              //creates a list of all possible combination of independent variables
	clean_combos := filteredExhaustiveList(combo_list) //removes any empty strings

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runConcur(n_times, clean_combos, dfMatrix, headers, mvSlice)
	}
}

func BenchmarkS(b *testing.B) {

	n_times := 100
	input_filename := "boston.csv"

	//do the prep work to read in and edit the csv
	//this replicates the first part of func main
	csvFile, _ := ReadCSV(input_filename)
	defer csvFile.Close()
	df := dataframe.ReadCSV(csvFile)
	df = createLevels(df, "neighborhood")
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
	mv := df.Select([]string{"mv"})
	df = df.Drop([]string{"mv"})
	mvRecords := mv.Records()
	mvSlice := depedentSlicer(mvRecords)
	records := df.Records()
	dfMatrix := indepdentSlicer(records)
	headers := df.Names()                              //grab all headers
	combo_list := exhaustiveList(headers)              //creates a list of all possible combination of independent variables
	clean_combos := filteredExhaustiveList(combo_list) //removes any empty strings

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runSeq(n_times, clean_combos, dfMatrix, headers, mvSlice)
	}
}
