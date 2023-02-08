/*
 Utilizing the golearn package to perform k-means clustering on a dataset. 
 The first step is to load the dataset into a base.Instances struct, 
 then a k-means clusterer is created using the cluster.NewKMeansClusterer function and perform clustering using the cluster.KMeans function. 
 The quality of the clustering is evaluated using the Silhouette score and
 the first five instances of each cluster is printed to verify that the data has been clustered correctly.

*/

package data_clustering


import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/cluster"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func DataClustering(csv_file_path string) {
	// Load the iris dataset.
	rawData, err := base.ParseCSVToInstances(csv_file_path, true)
	if err != nil {
		fmt.Println(err)
		return
	}
  
  
	// Perform k-means clustering on the data.
	kmeans := cluster.NewKMeansClusterer(3, distance.EuclideanDistance, 10)
	clusters, err := cluster.KMeans(kmeans, rawData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Evaluate the quality of the clustering.
	evaluator, err := evaluation.NewClusterEvaluator(rawData, clusters, distance.EuclideanDistance)
	if err != nil {
		fmt.Println(err)
		return
	}
	silhouette := evaluator.Silhouette()
	fmt.Printf("Silhouette score: %0.2f\n", silhouette)

	// Print the first five instances of each cluster.
	fmt.Println("First five instances of each cluster:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Cluster %d:\n", i)
		for j := 0; j < 5; j++ {
			fmt.Println(clusters[i].RowString(j))
		}
		fmt.Println("")
	}
}

/*
Presumed Output
Silhouette score: 0.55
First five instances of each cluster:
Cluster 0:
[5.1 3.5 1.4 0.2]
[4.9 3.  1.4 0.2]
[4.7 3.2 1.3 0.2]
[4.6 3.1 1.5 0.2]
[5.  3.6 1.4 0.2]

Cluster 1:
[5.1 3.5 1.4 0.2]
[4.9 3.  1.4 0.2]
[4.7 3.2 1.3 0.2]
[4.6 3.1 1.5 0.2]
[5.  3.6 1.4 0.2]

Cluster 2:
[5.1 3.5 1.4 0.2]
[4.9 3.  1.4 0.2]
[4.7 3.2 1.3 0.2]
[4.6 3.1 1.5 0.2]
[5.  3.6 1.4 0.2]

Cluster 3:
[5.1 3.5 1.4 0.2]
[4.9 3.  1.4 0.2]
[4.7 3.2 1.3 0.2]
[4.6 3.1 1.5 0.2]
[5.  3.6 1.4 0.2]

Cluster 4:
.....
*/
