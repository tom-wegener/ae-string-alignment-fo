package main

// x[i][j] is the flow for the between the vertices i and j
// The functions get the three cost-matrices (variable a, varioable b, fixed c) and the flow-matrix and calculate the overall cost of the flow-matrix.
// We can use them as to calculate our overall fitness

func (x *Child) costCalculator(a, b, c [][]int64, customerDemand []int64) {
	function := 0

	if function == 0 {
		x.costEstimatorZero(a, b, c, customerDemand)
	}
}

func (x *Child) costEstimatorZero(a, b, c [][]int64, customerDemand []int64) {
	var costsX int64
	var costs int64
	for i, row := range a {
		for j := range row {
			costsX = costsX + costEstimatorZeroEdge(a[i][j], b[i][j], c[i][j], x.flow[i][j])
		}
		costs = costs + costsX
	}
	x.fitness = costs
	demandChecker(x.flow, customerDemand)
	return
}

func costEstimatorZeroEdge(aij, bij, cij int64, xij int64) (costs int64) {
	costs = -aij*xij*xij + bij*xij + cij
	return

}

// TODO: Think about a smart way to model this shit...
func demandChecker(flow [][]int64, customerDemand []int64) (strafterm int64) {

	strafterm = 1
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
