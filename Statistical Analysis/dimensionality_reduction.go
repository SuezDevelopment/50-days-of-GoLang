/*
Utilizing the golearn package to perform dimensionality reduction on a dataset. 
The first step is to load the dataset into a base.Instances struct, 
then create a PCA (Principal Component Analysis) reducer using the pca.NewPCA function, 
and reduce the dimensionality of the data using the dimred.ApplyPCA function. 
Finally,the first five instances of the reduced data is printed to verify that the dimensionality has been reduced.
*/

package dimensionality_reduction

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/dimred"
	"github.com/sjwhitworth/golearn/pca"
)

func DimensionalityReduction(csv_file_path string) {
	// Load the iris dataset.
	rawData, err := base.ParseCSVToInstances(csv_file_path, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Reduce the dimensionality of the data.
	pca := pca.NewPCA(2)
	reducedData, err := dimred.ApplyPCA(pca, rawData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the first five instances of the reduced data.
	fmt.Println("First five instances of the reduced data:")
	for i := 0; i < 5; i++ {
		fmt.Println(reducedData.RowString(i))
	}
}

/*
Presumed Output
First five instances of the reduced data:
[5.1 3.5]
[4.9 3. ]
[4.7 3.2]
[4.6 3.1]
[5.  3.6]

*/
