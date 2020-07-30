package main

import (
	"math/rand"
	"time"
)

//Child is a child or genome with its fitness
type Child struct {
	flow    [][]int64
	demand  []int64
	storage []int64
	fitness int64
}

func (x *Child) toParent(c *Child) {
	c.flow = x.flow
	c.fitness = x.fitness
	c.storage = x.storage
	c.demand = x.demand
}

// mutateDumb: mutate the genome on all places but random place
//  - mutates all
func (x *Child) mutateDumb(c *Child) {
	c.flow = x.flow
	c.fitness = x.fitness
	c.storage = x.storage
	c.demand = x.demand
	localStorage := x.demand

	for i, row := range x.flow {
		for j := range row {
			randomInt := rand.Int63n(10)
			localStorage[i] = localStorage[i] - randomInt
			localStorage[j] = localStorage[j] + randomInt
			c.flow[i][j] = randomInt
		}
	}
}

// find a new Neighbour but still relatively random
//  - only mutates one
func (x *Child) findNeighbourZero(c *Child) {
	c.flow = x.flow
	c.fitness = x.fitness
	c.storage = x.storage
	c.demand = x.demand
	localStorage := x.demand

	randomI := rand.Int63n(10)
	randomJ := rand.Int63n(9)
	randomInt := rand.Int63n(10)

	c.flow[randomI][randomJ] = randomInt

	for i := randomI; i >= 0; i-- {
		for j := randomJ; j >= 0; j-- {
			localStorage[i] = localStorage[i] - x.flow[i][j]
			localStorage[j] = localStorage[j] + x.flow[i][j]
			c.flow[i][j] = randomInt
		}
	}
	c.storage = localStorage
	if c.flow[randomI][randomJ] != randomInt {

	}

}

// find a new Neighbour
//  - mutates one randomly
//  - corrects multiple after the mutated
//  - recognize demand/storage
func (x *Child) findNeighbourOne(c *Child) {
	c.flow = x.flow
	c.fitness = x.fitness
	c.storage = x.storage
	c.demand = x.demand
	localStorage := x.demand

	randomI := rand.Intn(10)

	randomJ := rand.Intn(9)

	for i := randomI; i >= 0; i-- {
		for j := randomJ; j >= 0; j-- {
			if localStorage[i] > 0 {
				randomInt := rand.Int63n(localStorage[i])
				localStorage[i] = localStorage[i] - randomInt
				localStorage[j] = localStorage[j] + randomInt
				c.flow[i][j] = randomInt
			}
		}
	}
	c.storage = localStorage
	return
}

// find a new Neighbour
//  - only mutates one
//  - recognize demand/storage
//  - only mutate existing edges
func (x *Child) findNeighbourTwo(c *Child, network [][]bool) {
	c.flow = x.flow
	c.fitness = 0
	c.storage = x.storage
	c.demand = x.demand
	localStorage := x.demand

	var edges [][]int
	for i, row := range network {
		for j, cell := range row {
			if cell {
				edge := []int{i, j}
				edges = append(edges, edge)
			}
		}
	}

	randomEdge := rand.Intn(len(edges))

	for i := len(c.flow) - 1; i >= edges[randomEdge][0]; i-- {
		for j := edges[randomEdge][1]; j < len(c.flow[i]); j++ {
			tmp := localStorage[i] + c.flow[i][j]
			if tmp > 0 {
				randomInt := rand.Int63n(tmp)
				localStorage[i] = localStorage[i] - randomInt
				localStorage[j] = localStorage[j] + randomInt
				c.flow[i][j] = randomInt
			}
		}
	}
	c.storage = localStorage
	return

}

// initiateFlowDumb: this function generates a first very random genome
func (x *Child) initiateFlowZero(verticesCount int) {

	var flowAll [][]int64

	for i := 0; i < verticesCount; i++ {
		var flowX []int64
		for j := 0; j < (verticesCount - 1); j++ {
			randomInt := rand.Int63n(10)
			flowX = append(flowX, randomInt)
		}
		flowAll = append(flowAll, flowX)
	}
	x.flow = flowAll
	return
}

// initiateFlowSmarter: this function generates a first flow
//  - only uses existing vertices
func (x *Child) initiateFlowOne(verticesCount int, network [][]bool) {

	localStorage := x.demand
	for i := (verticesCount - 1); i >= 0; i-- {
		var flowX []int64
		for j := (verticesCount - 2); j >= 0; j-- {
			if network[i][j] {
				randomInt := rand.Int63n(10)
				localStorage[i] = localStorage[i] - randomInt
				localStorage[j] = localStorage[j] + randomInt
				flowX = append(flowX, randomInt)
			} else {
				flowX = append(flowX, 0)
			}
		}
		x.flow = append(x.flow, flowX)
	}
	x.storage = localStorage
	return
}

// initiateFlowTwo: this function generates a first flow
//  - only uses existing vertices
//  - recognizes the demand
func (x *Child) initiateFlowTwo(verticesCount int, network [][]bool) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	localStorage := make([]int64, len(x.demand))
	copy(localStorage, x.demand)

	for i := (verticesCount - 1); i >= 0; i-- {
		var flowX []int64
		for j := (verticesCount - 2); j >= 0; j-- {
			if network[i][j] && localStorage[i] > 0 {
				randomInt := r.Int63n(localStorage[i])
				localStorage[i] = localStorage[i] - randomInt
				localStorage[j] = localStorage[j] + randomInt
				flowX = append(flowX, randomInt)
			} else {
				flowX = append(flowX, 0)
			}
		}
		x.flow = append(x.flow, flowX)
	}
	x.storage = localStorage
	return
}
