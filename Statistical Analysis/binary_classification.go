/*

 Here I used the golearn package to perform classification on a dataset. 
 The first step is to load the dataset into a base.Instances struct, 
 then split the data into a training set and a test set using the base.InstancesTrainTestSplit function. 
 Next,a KNN classifier is created using the knn.NewKnnClassifier function, 
 fit the classifier to the training data using the Fit method, and make predictions on the test data using the Predict method. 
 Finally, the evaluation.GetConfusionMatrix function is used to get a confusion matrix, which summarizes the performance of the classifier, 
 and the evaluation.GetSummary function to get a summary of the accuracy of the classifier.

*/

package binary_classification

import (
	"fmt"
	"math/rand"
   "encoding/csv"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func BinaryClassification(csvfilename string) {
	// Load the iris dataset.
	rawData, err := base.ParseCSVToInstances(csvfilename, true)
	if err != nil {
		panic(err)
	}

	// Split the data into training and test sets.
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.75)

	// Train the KNN classifier.
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)
	cls.Fit(trainData)

	// Evaluate the classifier on the test set.
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}
  
  // Print the accuracy of the classifier.
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(err)
	}

	// Print the confusion matrix.
	fmt.Println(evaluation.GetSummary(confusionMat))
}


/*
Presumed Output
Accuracy: 0.97
Kappa: 0.94
Confusion Matrix (Row=True, Col=Predicted):
 [[14  0  0]
 [ 0 12  1]
 [ 0  1 12]]
*/
