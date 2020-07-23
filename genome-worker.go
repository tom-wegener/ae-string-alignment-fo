package main

import (
	"crypto/rand"
	"math/big"
)

//Child is a child or genome with its fitness
type Child struct {
	flow    [][]int64
	storage []int64
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
func (x *Child) initiateFlowZero(verticesCount int) {

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
func (x *Child) initiateFlowOne(verticesCount int, customerDemand []int64, network [][]bool) {
	var flowAll [][]int64

	for i := verticesCount; i > 0; i-- {
		var flowX []int64
		for j := (verticesCount - 1); j > 0; j-- {
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

// initiateFlowSmarter: this function generates a first flow which only uses existing vertices
func (x *Child) initiateFlowTwo(verticesCount int, customerDemand []int64, network [][]bool, sourceCapacity int64) {
	var flowAll [][]int64

	// The demand could be seen as negative storage capacity
	// Also the last node is the source
	for i := range customerDemand {
		x.storage = append(x.storage, -1*customerDemand[i])
	}
	x.storage = append(x.storage, sourceCapacity)

	for i := (verticesCount - 1); i >= 0; i-- {
		var flowX []int64
		for j := (verticesCount - 2); j >= 0; j-- {
			if network[i][j] && x.storage[i] > 0 {
				randomInt, _ := rand.Int(rand.Reader, big.NewInt(x.storage[i]))
				x.storage[i] = x.storage[i] - randomInt.Int64()
				x.storage[j] = x.storage[j] + randomInt.Int64()
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
