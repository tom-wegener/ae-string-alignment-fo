package main

import "math"

// x[i][j] is the flow for the between the vertices i and j
// The functions get the three cost-matrices (variable a, varioable b, fixed c) and the flow-matrix and calculate the overall cost of the flow-matrix.
// We can use them as to calculate our overall fitness

func (x *Child) costCalculator(a, b, c [][]int64) {
	function := 0

	if function == 0 {
		x.costEstimatorZero(a, b, c)
	}
}

func (x *Child) costEstimatorZero(a, b, c [][]int64) {
	var costsX int64
	var costs int64
	for i, row := range a {
		for j := range row {
			costsX = costsX + costEstimatorZeroEdge(a[i][j], b[i][j], c[i][j], x.flow[i][j])
		}
		costs = costs + costsX
	}
	var punishment float64
	for _, store := range x.storage {
		storeFloat := float64(store)
		punishment = punishment + math.Abs(storeFloat)
	}
	punishmentInt := int64(punishment) * 100000
	x.fitness = costs + punishmentInt
	return
}

func costEstimatorZeroEdge(aij, bij, cij int64, xij int64) (costs int64) {
	costs = -aij*xij*xij + bij*xij + cij
	return

}
