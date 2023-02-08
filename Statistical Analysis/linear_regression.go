package linear_regression
/*
Here there's a reuseable function that takes two arrays xData and yData that represent the independent and dependent variables of our data. 
Then used the stat.LinearRegression function to fit a linear regression model to the data. 
The function returns the intercept and slope of the best-fit line to make predictions for new values of x. 
In this example,x = 6.0 and using the fitted model to predict the corresponding value of y.
*/

func LinearRegression(xData []float64, yData []float64) {
	// Fit the linear regression model.
	var intercept, slope float64
	intercept, slope = stat.LinearRegression(xData, yData, nil, false)

	// Print the results.
	fmt.Printf("Intercept: %0.2f\n", intercept)
	fmt.Printf("Slope: %0.2f\n", slope)

	// Predict the value of y for a given value of x.
	x := 6.0
	y := intercept + slope*x
	fmt.Printf("Predicted value of y for x = %0.2f: %0.2f\n", x, y)
}


xData := []float64{1, 2, 3, 4, 5}
yData := []float64{2, 4, 6, 8, 10}

LinearRegression(xData, yData) 
/* Output Intercept: 0.00 
Slope: 2.00
Predicted value of y for x = 6.00: 12.00
*/




