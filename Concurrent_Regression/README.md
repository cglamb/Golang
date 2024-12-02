# Go-Based Linear Regression with Concurrency

- **Author**: Charles Lamb
- **Contact Info**: charlamb@gmail.com
- **Github address for this project**: [https://github.com/cglamb/Concurrent_Regression](https://github.com/cglamb/Concurrent_Regression)
- **Git Clone command for the repository**: `git clone https://github.com/cglamb/Concurrent_Regression.git`

## Introduction

This project applies a number of linear models to Boston housing data. The underlying dataset contains 13 independent variables and 1 dependent variable. All combinations of independent variables are considered. An optimal model is selected based on the lowest Akaike Information Criterion (AIC).

The application is built to test performance differentials when running linear models concurrently versus sequentially. As such, the application is built to allow the user to run the regression modeling using either sequential or concurrent processing.

## Findings

Concurrency was found to substantially reduce the runtime required to complete the tested models. 819,100 regression models were run concurrently in 13.498s in this application. Running sequentially, the same number of models took 53.988s to complete, implying concurrency reduced runtime down to 1/4th of the time required to run sequentially. These findings suggest that it could be potentially highly beneficial for organizations to explore applications that take advantage of concurrency, particularly for heavy/long computational tasks.

## Explanation of Files

- `Logs/`
  - `Benchmark_log` – log of benchmark testing
  - `Executable_log` – log of .exe being run
  - `Testing_log` – log of testing applications in `main_test.go` being run
- `Boston.csv` – underlying data
- `Csv_utilities.go` – contains functions used to manipulate CSV files
- `Data_cleaning.go` – contains functions used to prep data for regression analysis
- `Exhaustive_regression.exe` – compiled executable function
- `Exhaustive_utility.go` – functions used to find all possible combinations of independent variables
- `File_utility.go` – library for future use
- `Go.mod` – golang dependency file
- `Go.sum` – Golang dependency file
- `Main.go` – performs regression analysis
- `Regression.go` – contains functions used to run regression and calculate relevant statistics
- `Test_cases.csv` – file used for `main_test.go`

## Running the Command Line Application

If the terminal's current directory is the directory containing the executable, the program can be run from the command line using the following command: `./exhaustive_regression -i number -c=true/false`. Number can be specified as any integer value greater than 0. This is a duplication variable used for benchmarking, and instructs the application to run each regression analysis a number of times equivalent to `number`. `-c` is a Boolean switch that determines where the application runs the regression concurrently or sequentially. `True` results in concurrent processing.

## Input Data

The data to which the regression models are applied is the `Boston.CSV` contained in the repository. The data consists of 13 independent variables and 1 dependent variable. The “mv” variable is the dependent variable. The following map describes each variable in greater detail.

| Variable  | Description                                           |
|-----------|-------------------------------------------------------|
| Neighborhood | Name of the Boston neighborhood (census track location) |
| mvMedian  | Value of homes in thousands of 1970 US dollars       |
| nox       | Air pollution (nitrogen oxide concentration)         |
| crim      | Crime rate                                            |
| zn        | Percentage of land zoned for lots                    |
| indus     | Percentage of business that is industrial or nonretail |
| chas      | On the Charles River (1) or not (0)                  |
| rooms     | Average number of rooms per home                     |
| age       | Percentage of homes built before 1940                |
| dis       | Weighted distance to employment centers              |
| rad       | Accessibility to radial highways                     |
| tax       | Tax rate                                             |
| ptratio   | Pupil/teacher ratio in public schools                |
| lstat     | Percentage of population of lower socio-economic status |

**Data source**: Miller, Thomas W. 1999. "The Boston splits: Sample size requirements for modern regression." 1999 Proceedings of the Statistical Computing Section of the American Statistical Association, 210–215.

## Compiling Instructions

The executable can be compiled in Go via the `go build` command within the terminal.

## Testing

Testing functions were developed for a variety of situations and can be located in `main_test.go`.

## Enhancements

Some of the data manipulation performed by the application is specific to the `boston.csv` data. This includes hard coding of fields that need to be leveled, hard coding of float conversions, and hard coding of identification of the dependent variable. The application could be
