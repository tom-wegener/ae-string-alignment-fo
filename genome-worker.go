package main

import (
	"crypto/rand"
	"math/big"
)

//Child is a child or genome with its fitness
type Child struct {
	flow    [][]int64
	fitness int64
}

func (x *Child) toParent(c *Child) {
	c.flow = x.flow
	c.fitness = x.fitness
}

// mutateDumb: mutate the genome on a random place
func (x *Child) mutateDumb() {

	for i, row := range x.flow {
		for j := range row {
			randomInt, err := rand.Int(rand.Reader, big.NewInt(10))
			errFunc(err)
			if x.flow[i][j] != randomInt.Int64() {
				x.flow[i][j] = randomInt.Int64()
			}
		}
	}
}

//
func (x *Child) findNeighbour(c *Child) {
	c.flow = x.flow
	c.fitness = x.fitness
	randomI, err := rand.Int(rand.Reader, big.NewInt(10))
	errFunc(err)
	randomI64 := randomI.Int64()
	randomJ, err := rand.Int(rand.Reader, big.NewInt(9))
	errFunc(err)
	randomJ64 := randomJ.Int64()
	randomInt, err := rand.Int(rand.Reader, big.NewInt(10))
	if c.flow[randomI64][randomJ64] != randomInt.Int64() {
		c.flow[randomI64][randomJ64] = randomInt.Int64()
	}

}

// initiateFlowDumb: this function generates a first very random genome
func (x *Child) initiateFlowDumb(verticesCount int) {

	var flowAll [][]int64

	for i := 0; i < verticesCount; i++ {
		var flowX []int64
		for j := 0; j < (verticesCount - 1); j++ {
			randomInt, _ := rand.Int(rand.Reader, big.NewInt(10))
			flowX = append(flowX, randomInt.Int64())
		}
		flowAll = append(flowAll, flowX)
	}
	x.flow = flowAll
	return
}

// initiateFlowSmarter: this function generates a first flow which only uses existing vertices
func (x *Child) initiateFlowSmarter(verticesCount int, customerDemand []int64, network [][]bool) {
	var flowAll [][]int64

	for i := 0; i < verticesCount; i++ {
		var flowX []int64
		for j := 0; j < (verticesCount - 1); j++ {
			if network[i][j] {
				randomInt, _ := rand.Int(rand.Reader, big.NewInt(10))
				flowX = append(flowX, randomInt.Int64())
			} else {
				flowX = append(flowX, 0)
			}
		}
		flowAll = append(flowAll, flowX)
	}
	x.flow = flowAll
	return
}
