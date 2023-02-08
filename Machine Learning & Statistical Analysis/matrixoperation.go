/*
Utilizing the gorgonia package to perform matrix multiplication. 
First, two matrices m1 and m2 are created using the tensor package. 
Next, a gorgonia.Graph is created to represent the computation graph, 
then nodes for the matrices are created using gorgonia.NewMatrix. 
To perform the matrix multiplication, the gorgonia.Mul function is used. 
A gorgonia.TapeMachine is then created to execute the graph and executed using machine.RunAll. 
Finally, the result is retrieved using resultNode.Value() and printed.
*/

package matrix_operation

import (
	"fmt"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func MultiplyMatrix(m1 [][]int, m2 [][]int) [][]int {
	// Create a Gorgonia graph.
	g := gorgonia.NewGraph()
	/
  / Create nodes for the matrices.
	m1Node := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(2, 2), gorgonia.WithValue(m1))
	m2Node := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(2, 2), gorgonia.WithValue(m2))

	// Perform matrix multiplication.
	resultNode, err := gorgonia.Mul(m1Node, m2Node)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a machine to execute the graph.
	machine := gorgonia.NewTapeMachine(g)

	// Run the machine.
	if err = machine.RunAll(); err != nil {
		fmt.Println(err)
		return
	}

	// Get the result.
	resultVal, err := resultNode.Value().(tensor.Tensor).Data()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the result.
	fmt.Println("Result:")
	fmt.Println(resultVal)
}

// Create two matrices.
m1 := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float64{1, 2, 3, 4}))
m2 := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float64{5, 6, 7, 8}))

MultiplyMatrix(m1,m2)
/*
Presumed Output
Result:
[19. 22. 43. 50.]
*/
