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

// mutate: mutate the genome on a random place
func (c *Child) mutate() {

	for i, row := range c.flow {
		for j := range row {
			randomInt, _ := rand.Int(rand.Reader, big.NewInt(10))
			if c.flow[i][j] != randomInt.Int64() {
				c.flow[i][j] = randomInt.Int64()
			}
		}
	}
	println(c.flow[5][5])
}

// initiateFlow: this function generates a first genome
func (c *Child) initiateFlow(verticesCount int) {
	var flowX []int64
	var flowAll [][]int64

	for i := 0; i < verticesCount; i++ {
		for j := 0; j < (verticesCount - 1); j++ {
			randomInt, _ := rand.Int(rand.Reader, big.NewInt(10))
			flowX = append(flowX, randomInt.Int64())
		}
		flowAll = append(flowAll, flowX)
	}
	c.flow = flowAll
	return
}
