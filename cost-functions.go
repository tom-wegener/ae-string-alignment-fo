package main

// x[i][j] is the flow for the between the vertices i and j
// The functions get the three cost-matrices (variable a, varioable b, fixed c) and the flow-matrix and calculate the overall cost of the flow-matrix.
// We can use them as to calculate our overall fitness

func (x *Child) costEstimatorOne(a, b, c [][]int64, verticesCount int) {
	println(x.flow[5][5])
	var costsX int64
	var costs int64
	for i := 0; i < verticesCount; i++ {
		for j := 0; j < (verticesCount - 1); j++ {
			costsX = costsX + costEstimatorOneEdge(a[i][j], b[i][j], c[i][j], x.flow[i][j])
		}
		costs = costs + costsX
	}
	x.fitness = costs
	return

}

func costEstimatorOneEdge(aij, bij, cij int64, xij int64) (costs int64) {
	costs = -aij*xij*xij + bij*xij + cij
	return

}

/*
func costEstimatorTwo(a, b, c, x [][]int, verticesCount int) (costs int){
	var costsX int
	for i := 0; i < verticesCount; i++ {
		for j := 0; j < verticesCount; j++ {
			costsX = costsX + (-a[i][j]*x[i][j]*x[i][j] + b[i][j]*x[i][j] + c[i][j])
		}
		costs = costs + costsX
	}
	return
}

func costEstimatorThree(a, b, c, x [][]int, verticesCount int) (costs int){
	var costsX int
	for i := 0; i < verticesCount; i++ {
		for j := 0; j < verticesCount; j++ {
			costsX = costsX + (-a[i][j]*x[i][j]*x[i][j] + b[i][j]*x[i][j] + c[i][j])
		}
		costs = costs + costsX
	}
	return
}
*/
