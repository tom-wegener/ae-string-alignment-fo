package main

import "os"

// x[i][j] is the flow for the between the vertices i and j
// The functions get the three cost-matrices (variable a, varioable b, fixed c) and the flow-matrix and calculate the overall cost of the flow-matrix.
// We can use them as to calculate our overall fitness

func (x *Child) costCalculator(a, b, c [][]int64) {
	var costs int64 = 0
	for i, row := range a {
		var costsX int64
		for j := range row {
			costsX = costsX + costCalculatorEdge(a[i][j], b[i][j], c[i][j], x.flow[i][j])
		}
		costs = costs + costsX
	}

	var punishment int64 = 0
	for _, store := range x.storage {
		if store != 0 {
			punishment = punishment + abs(store)
		}
	}
	costs = costs + punishment*100
	x.fitness = costs
	if x.fitness < 0 {
		println(x.fitness)
		for i, row := range x.flow {
			print(x.storage[i], ", ")
			for j := range row {
				print(x.flow[i][j], ", ")
			}
			println()
		}
		os.Exit(0)
	}
}

func costCalculatorEdge(aij, bij, cij int64, xij int64) (costs int64) {
	costs = aij*xij*xij + bij*xij + cij
	if costs < 0 {
		println(-aij, xij, bij, xij, cij)
	}
	return

}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}
