package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// meanSquareError calculates the mean square error between the actual and predicted values.
func meanSquareError(actual, predicted []float64) float64 {
	var sum float64 //initalize variable to hold SSE
	//calculate the sum of the square error
	for i := range actual {
		diff := actual[i] - predicted[i]
		sum += diff * diff
	}
	//calculate the mse
	mse := sum / float64(len(actual))
	return mse
}

// fit a least squares regresion line
// pass to the function indepedent variables and dependent variables
// functions exports the linear regression fitted coefficients
// i couldnt find a library that handles OLS for multi-variable cases, so took the approach of doing the matrix alegebra.
// https://max.pm/posts/ols_matrix/
func LinearRegressionFit(X [][]float64, Y []float64) ([]float64, error) {

	//putting the independent variable data into a matrix
	rows, cols := len(X), len(X[0])+1        //need an extra column for beta0 intercept
	Xmatrix := mat.NewDense(rows, cols, nil) //initiazling the matrix
	for i, xi := range X {                   //feeding the slice of slice data from X into the matrix
		Xmatrix.SetRow(i, append([]float64{1}, xi...))
	}

	// seeting up the matrices we will be calculating
	var XprimtX mat.Dense
	var XprimeXinverse mat.Dense
	var coeffmatrix mat.Dense
	var XTY mat.Dense

	//setting up the depedent variable in a matrix
	Ymatrix := mat.NewDense(len(Y), 1, Y)

	// dot product of x prime and x
	XprimtX.Mul(Xmatrix.T(), Xmatrix) //product of the xprime and x matrix

	// calculate the matrix inverse
	err := XprimeXinverse.Inverse(&XprimtX) //apply the inverse function
	if err != nil {
		return nil, fmt.Errorf("error in calculating matrix inverse.  see LinearRegressionFit Function.  problem at XprimeXinverse")
	}

	// calculating transpose(x) * y
	XTY.Mul(Xmatrix.T(), Ymatrix)

	// computing the coefficient matrix
	//coefficient matrix is inverse[ dot product of X transpose and X] * transpose(x)y
	coeffmatrix.Mul(&XprimeXinverse, &XTY)

	// converting the coefficent matrix into a slice
	beta := make([]float64, cols)
	for i := 0; i < cols; i++ {
		beta[i] = coeffmatrix.At(i, 0)
	}

	return beta, nil
}

// generate yhat given observed xs and the linear regression coefficients
// function predicts y values expected by an OLS regression
func Predict(x [][]float64, beta []float64) []float64 {

	cnt := len(x)                //number of prediction to make
	yhat := make([]float64, cnt) //initializing the prediction slice.  has length equal to number of passed x values

	//looping through each x value
	//for each x calculating a yhat
	for i, x := range x {

		// Include the intercept in the prediction calculation
		pred := beta[0] // Start with the intercept
		for j, xi := range x {
			pred += beta[j+1] * xi // Add the product of each coefficient and observed value
		}
		yhat[i] = pred
	}

	return yhat
}
